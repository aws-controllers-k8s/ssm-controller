package document

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	ackmetrics "github.com/aws-controllers-k8s/runtime/pkg/metrics"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	svcsdk "github.com/aws/aws-sdk-go-v2/service/ssm"

	svcapitypes "github.com/aws-controllers-k8s/ssm-controller/apis/v1alpha1"
)

func TestSDKFindReadsObservedDocumentContent(t *testing.T) {
	const (
		name           = "test-document"
		desiredContent = "schemaVersion: '2.2'\ndescription: desired\n"
		awsContent     = "schemaVersion: '2.2'\ndescription: observed\n"
	)

	var operations []string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		operation := r.Header.Get("X-Amz-Target")
		switch {
		case strings.HasSuffix(operation, ".DescribeDocument"):
			operations = append(operations, "DescribeDocument")
			w.Header().Set("Content-Type", "application/x-amz-json-1.1")
			_, _ = w.Write([]byte(`{"Document":{"Name":"test-document","DocumentVersion":"1","DocumentFormat":"YAML","DocumentType":"Command","Status":"Active","LatestVersion":"1","DefaultVersion":"1"}}`))
		case strings.HasSuffix(operation, ".GetDocument"):
			operations = append(operations, "GetDocument")
			var input struct {
				Name            string `json:"Name"`
				DocumentVersion string `json:"DocumentVersion"`
				DocumentFormat  string `json:"DocumentFormat"`
			}
			if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
				t.Errorf("decoding GetDocument input: %v", err)
			}
			if input.Name != name || input.DocumentVersion != "1" || input.DocumentFormat != "YAML" {
				t.Errorf("unexpected GetDocument input: %+v", input)
			}
			w.Header().Set("Content-Type", "application/x-amz-json-1.1")
			_, _ = w.Write([]byte(`{"Name":"test-document","DocumentVersion":"1","DocumentFormat":"YAML","Content":"schemaVersion: '2.2'\ndescription: observed\n"}`))
		case strings.HasSuffix(operation, ".UpdateDocument"):
			operations = append(operations, "UpdateDocument")
			var input struct {
				Name            string `json:"Name"`
				Content         string `json:"Content"`
				DocumentVersion string `json:"DocumentVersion"`
			}
			if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
				t.Errorf("decoding UpdateDocument input: %v", err)
			}
			if input.Name != name || input.Content != desiredContent || input.DocumentVersion != "1" {
				t.Errorf("unexpected UpdateDocument input: %+v", input)
			}
			w.Header().Set("Content-Type", "application/x-amz-json-1.1")
			_, _ = w.Write([]byte(`{"DocumentDescription":{"Name":"test-document","DocumentVersion":"2","DocumentFormat":"YAML","DocumentType":"Command","Status":"Active","LatestVersion":"2","DefaultVersion":"1"}}`))
		default:
			t.Errorf("unexpected SSM operation %q", operation)
			http.Error(w, "unexpected operation", http.StatusBadRequest)
		}
	}))
	defer server.Close()

	sdkapi := svcsdk.NewFromConfig(aws.Config{
		Region:      "us-east-1",
		Credentials: aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider("test", "test", "")),
	}, func(options *svcsdk.Options) {
		options.BaseEndpoint = aws.String(server.URL)
	})
	rm := &resourceManager{
		sdkapi:  sdkapi,
		metrics: ackmetrics.NewMetrics("ssm"),
	}
	desired := &resource{ko: &svcapitypes.Document{
		Spec: svcapitypes.DocumentSpec{
			Name:           aws.String(name),
			DocumentFormat: aws.String("YAML"),
			Content:        aws.String(desiredContent),
		},
		Status: svcapitypes.DocumentStatus{
			DocumentVersion: aws.String("1"),
		},
	}}

	latest, err := rm.sdkFind(context.Background(), desired)
	if err != nil {
		t.Fatalf("sdkFind returned error: %v", err)
	}
	if latest.ko.Spec.Content == nil || *latest.ko.Spec.Content != awsContent {
		t.Fatalf("sdkFind content = %q, want AWS content %q", aws.ToString(latest.ko.Spec.Content), awsContent)
	}
	delta := newResourceDelta(desired, latest)
	if !delta.DifferentAt("Spec.Content") {
		t.Fatalf("content delta was not detected: %+v", delta.Differences)
	}
	updated, err := rm.sdkUpdate(context.Background(), desired, latest, delta)
	if err != nil {
		t.Fatalf("sdkUpdate returned error: %v", err)
	}
	if updated.ko.Status.DocumentVersion == nil || *updated.ko.Status.DocumentVersion != "2" {
		t.Fatalf("updated document version = %q, want 2", aws.ToString(updated.ko.Status.DocumentVersion))
	}
	if want := []string{"DescribeDocument", "GetDocument", "UpdateDocument"}; !reflect.DeepEqual(operations, want) {
		t.Fatalf("SSM operations = %v, want %v", operations, want)
	}
}
