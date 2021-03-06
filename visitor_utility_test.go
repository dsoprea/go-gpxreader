package gpxreader

import (
    "testing"
    "bytes"

    "github.com/dsoprea/go-logging"
)

func TestEnumerateTrackPoints(t *testing.T) {
    n := 0
    cb := func(tp *TrackPoint) error {
        n++

        return nil
    }

    b := bytes.NewBufferString(TestGpxData)

    if err := EnumerateTrackPoints(b, cb); err != nil {
        log.Panic(err)
    }

    if n != 204 {
        t.Fatalf("Point count not correct: (%d)", n)
    }
}

func TestExtractTrackPoints(t *testing.T) {
    b := bytes.NewBufferString(TestGpxData)
    points, err := ExtractTrackPoints(b)
    log.PanicIf(err)

    if len(points) != 204 {
        t.Fatalf("Point count not correct: (%d)", len(points))
    }
}
