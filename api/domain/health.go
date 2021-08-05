package domain

import (
	"context"
)

type HealthInfo struct {
	FaceComparison struct {
		QueryCount int `json:"queryCount"`
		QPS        int `json:"qps"`
		Instances  []struct {
			SvcType           string `json:"svcType"`
			InstanceID        string `json:"instanceId"`
			Status            string `json:"status"`
			UpTime            int    `json:"upTime"`
			LastRequestStatus string `json:"lastRequestStatus"`
			LastRequestTime   int64  `json:"lastRequestTime"`
		} `json:"instances"`
	} `json:"face-comparison"`
	ThaiID struct {
		QueryCount int `json:"queryCount"`
		QPS        int `json:"qps"`
		Instances  []struct {
			SvcType           string `json:"svcType"`
			InstanceID        string `json:"instanceId"`
			Status            string `json:"status"`
			UpTime            int    `json:"upTime"`
			LastRequestStatus string `json:"lastRequestStatus"`
			LastRequestTime   int64  `json:"lastRequestTime"`
		} `json:"instances"`
	} `json:"thai_id"`
	Antispoofing struct {
		QueryCount int `json:"queryCount"`
		QPS        int `json:"qps"`
		Instances  []struct {
			SvcType           string `json:"svcType"`
			InstanceID        string `json:"instanceId"`
			Status            string `json:"status"`
			UpTime            int    `json:"upTime"`
			LastRequestStatus string `json:"lastRequestStatus"`
			LastRequestTime   int64  `json:"lastRequestTime"`
		} `json:"instances"`
	} `json:"antispoofing"`
}

type HealthUsecase interface {
	GetPrometheusHealthInfo(ctx context.Context) (string, error)
}

type HealthRepository interface {
	GetHealthInfo(ctx context.Context) (*HealthInfo, error)
}
