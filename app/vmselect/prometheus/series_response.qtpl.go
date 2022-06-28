// Code generated by qtc from "series_response.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line app/vmselect/prometheus/series_response.qtpl:1
package prometheus

//line app/vmselect/prometheus/series_response.qtpl:1
import (
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/querytracer"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/storage"
)

// SeriesResponse generates response for /api/v1/series.See https://prometheus.io/docs/prometheus/latest/querying/api/#finding-series-by-label-matchers

//line app/vmselect/prometheus/series_response.qtpl:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line app/vmselect/prometheus/series_response.qtpl:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line app/vmselect/prometheus/series_response.qtpl:9
func StreamSeriesResponse(qw422016 *qt422016.Writer, isPartial bool, mns []storage.MetricName, qt *querytracer.Tracer, qtDone func()) {
//line app/vmselect/prometheus/series_response.qtpl:9
	qw422016.N().S(`{"status":"success","isPartial":`)
//line app/vmselect/prometheus/series_response.qtpl:12
	if isPartial {
//line app/vmselect/prometheus/series_response.qtpl:12
		qw422016.N().S(`true`)
//line app/vmselect/prometheus/series_response.qtpl:12
	} else {
//line app/vmselect/prometheus/series_response.qtpl:12
		qw422016.N().S(`false`)
//line app/vmselect/prometheus/series_response.qtpl:12
	}
//line app/vmselect/prometheus/series_response.qtpl:12
	qw422016.N().S(`,"data":[`)
//line app/vmselect/prometheus/series_response.qtpl:14
	for i := range mns {
//line app/vmselect/prometheus/series_response.qtpl:15
		streammetricNameObject(qw422016, &mns[i])
//line app/vmselect/prometheus/series_response.qtpl:16
		if i+1 < len(mns) {
//line app/vmselect/prometheus/series_response.qtpl:16
			qw422016.N().S(`,`)
//line app/vmselect/prometheus/series_response.qtpl:16
		}
//line app/vmselect/prometheus/series_response.qtpl:17
	}
//line app/vmselect/prometheus/series_response.qtpl:17
	qw422016.N().S(`]`)
//line app/vmselect/prometheus/series_response.qtpl:20
	qt.Printf("generate response: series=%d", len(mns))
	qtDone()

//line app/vmselect/prometheus/series_response.qtpl:23
	streamdumpQueryTrace(qw422016, qt)
//line app/vmselect/prometheus/series_response.qtpl:23
	qw422016.N().S(`}`)
//line app/vmselect/prometheus/series_response.qtpl:25
}

//line app/vmselect/prometheus/series_response.qtpl:25
func WriteSeriesResponse(qq422016 qtio422016.Writer, isPartial bool, mns []storage.MetricName, qt *querytracer.Tracer, qtDone func()) {
//line app/vmselect/prometheus/series_response.qtpl:25
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/prometheus/series_response.qtpl:25
	StreamSeriesResponse(qw422016, isPartial, mns, qt, qtDone)
//line app/vmselect/prometheus/series_response.qtpl:25
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/prometheus/series_response.qtpl:25
}

//line app/vmselect/prometheus/series_response.qtpl:25
func SeriesResponse(isPartial bool, mns []storage.MetricName, qt *querytracer.Tracer, qtDone func()) string {
//line app/vmselect/prometheus/series_response.qtpl:25
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/prometheus/series_response.qtpl:25
	WriteSeriesResponse(qb422016, isPartial, mns, qt, qtDone)
//line app/vmselect/prometheus/series_response.qtpl:25
	qs422016 := string(qb422016.B)
//line app/vmselect/prometheus/series_response.qtpl:25
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/prometheus/series_response.qtpl:25
	return qs422016
//line app/vmselect/prometheus/series_response.qtpl:25
}
