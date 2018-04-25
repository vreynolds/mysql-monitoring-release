package metrics_test

import (
	. "mysql-metrics/metrics"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Metrics", func() {
	It("is able to serialize a metric", func() {
		metric := Metric{Key: "key", Value: 3.14159, Unit: "unit"}
		Expect(metric.Key).To(Equal("key"))
		Expect(metric.Value).To(Equal(3.14159))
		Expect(metric.Unit).To(Equal("unit"))
	})
})
