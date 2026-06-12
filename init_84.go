package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "2.2" )

func lPsZwusbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Vm0Y5uliWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HSEpdoOQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KAJjZcd5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8NffnM8BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1LNVr1keWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wzWYnwsbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D2onns5UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v2vJzKJgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CA1zAquPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nIw27PTEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YvRZEANLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FtVcCvnUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wr1wwGLJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PxsieyorWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gVLae2qWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CLDHEIraWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NedqJSZ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cp5NgrSjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wDEHnUMkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V799jaHEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 739SFg6zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bgO6CMt0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YhSNrHe2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Vr3LuxBOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cJfqhZfXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U5r94C2xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NYveqRYxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IridsIHkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LVA8hM2HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ic7A4yjCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N9kgcOtsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JFtDSYg2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KIDplaNGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 76Cl4utpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FybVdlE7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OfZhATgcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nIAsA2baWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8ZBOlHnzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YiAseAhQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VJo9HRshWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 14wbe2oSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TYnmt1D5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Aa2EaAnmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func har1n2QwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FSPo9mrIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ypDJZC0vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hEtgTMTnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O9FK8w5bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Eo5jvtGKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6sA7RT8jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k4BidW4HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Jd7Tmk8eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Xymgz26zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PjSCl7lOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G3DdHfwpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Qfd0ChMBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vJM4MbSmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h0tBRWPMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fS25KJvoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bzuJE1IDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hS77y2P0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1JkqkNDeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TtdN9YI8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ip8oVlOoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bsjJUKTxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wXpqjIb2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fx7Dgr0CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LBwU1aDjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func swDPHemXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1BWDtk6rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8gnO2hUMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yM7sMV0UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func whdVT41JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P0K9l1FxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 69hqvaKNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uuNzlgUuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zaOXGaMfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hw8bUliHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bMtLCZtrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Xg6JyCRxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func szfdws68Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q2FypkeJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D2fX6DRdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4VLhHLiiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yET8PM90Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2QAf8QWsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NV87as2zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m05ScUBGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0U80QYQrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6LzegI99Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DR6JpZaGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BiF8tVhbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jTYsWodqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TdyqnvoBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iMRIaNMLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4LPMFRJpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D60fmojRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bePS5aAaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mqtNC1PxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BxKQemSmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k48c4Y7QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1RJ4P7tGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lIKHdkRpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zKfgOaskWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pSCcX6ivWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cdD4S3wBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r1DACD75Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TZLnlm2EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xevoWrPvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B5BW5qj2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JqezXhUiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WvJ0RmPUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CbIcjAlXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Qn2GsU90Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nbshhNnOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P1Rb8bwDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rvFIqusqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FN8jCKbnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tiKkXHiSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vbdv9WhLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S0TWvlykWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CtGrbpQwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7e19kpD4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RsK0ZbMHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gR03LylvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LWP8Ucw0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JkYwVn9HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2j8bkCTcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uWCv5s2bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XlNuHKKAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q2zNdPz1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yNttDYuMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 00HWAUXrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k5GHcRMhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yKCwLDXdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eavlbtfBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MRUmRYNSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mjtZa2Z8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IIIeEXJpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2IiJJWfHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8klmclQSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Qo75eylvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sT3KUNtuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IMcB1EzxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w9AqkC8lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jUa5OBpDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uSiEW1mOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pN1emuWmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aTXi743BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1fx7glggWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i96IEAxPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oOZvQhtKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HnaETiEqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t82WlTSiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hNCjuOlwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CGF3u8RaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZqXZ7vy0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KBMzlZpsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ptx2uvl0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uG3eqchpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VcpXznxaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XXeel78JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R7dqA0fSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func svvWx2koWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2zLGP3CJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LBtFO1JcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XIVzu6e5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zXgMMmLrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KrWMlaHkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y08X9J2PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kCBFFY44Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RjBYct54Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MMaBhshDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2ksVqJ5bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gbeFlfrNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7efYg76VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GkxMsFCHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hd9rZbzuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XF2oBqI4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3U6cyvmHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3iZdjzT9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1ABEBol5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HUGxVYjuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DlO4IVGdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y1u9zEXxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QL0Q9NrrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mmeARaHlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dnz36a8nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vwwKjg7iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N9El0rF6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WmWXRl9FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9X5oEJoxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yCnZPEBgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w9kcJ8QPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0en7A3M3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GNWIPqtMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5QRc0iMUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CQUkT1gVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7qiEvK66Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZBjch3HLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hzIMONKEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XtGhvVi6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wrZSgZ3CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QOjA2gaSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func syi3uxnYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5KZd6AxuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1uiLRrVkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H3VCdqVfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zqKznWQ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func upTcYVD1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nI1oievRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QBA5R1e9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func etqxZO5sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func awdCZ17QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XPHuyzRVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nM3GIX7wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func knzd6rC2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qUfacL3LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZPthJcXFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pxZviFBaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RgGaxPYpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wYmYgoNVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gMhIsW9rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func caajwA9LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XfSazzozWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EZiw6z20Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gn5fD3A4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T4UjXdC2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k9GEbM3dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MtyHKhZ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wIqcKkZUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func crIiHpqBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func knWp2bM1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ivf1FEiqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LsZNHInSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MwvMPGTTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3y9Nf9g3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xbIhMwztWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UxYnqoVZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Jp6PDDZsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qy14JVUOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MdDUMO58Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aCCW8ONEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OTscIeLuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func axgo2FWLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xIRBPhYvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H4Y77VcWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N7Tn7xyWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rtNTTQaGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OOXABhjRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QFNfOR5kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xyqCI5IhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qqMkYahsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gqdZi2CSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7ASMmUoJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rIMobXtpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7Gu84cQ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uB2ry94xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vHHADd0jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SFcAtsJHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g93i1a6RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Hedb9ZUgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func behlqt9mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vMnHzotuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1SHSEuLaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cqi2C4vqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dQi8crZ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d8c9mTEPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W9KhVlu8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yIOrQPUdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8h8uur4rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o6VZdvkjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fQwDTWP5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SGnAprHqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k3tv1iMlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PqgYGlPRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vj1zu7ZUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ig6AKkYqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PZPyIPQcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Skz65g5tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XIkfQiybWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1V1V9mwoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gpYfuCuuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ptFaaSMrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Rthuw9fTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iOG9FVYrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R0bEIPd9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SBl21eFjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ed4hryRMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W6Ie14yhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SoiS4cA0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IdidsbALWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qpuQkGQOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PZLBFSTTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4LHk3hjuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jx60kTh4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Hgp5b8neWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GlPluUrSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cyO5dbrNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gNeLnqj8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6lHdOyM7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QYAbsaqmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0qG7jeeoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fVN1GNz1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mwdrThqhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IzJLGrpdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UPDMDcIyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Zo2fCPOxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func smK22heZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4yW6tBOqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1KuEoSYqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Op0K4dSRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KQm0LZeoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TRLmW8r8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1BSx5YiEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EiVCLPV8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NgaKFQCBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1W8CbxxAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HVva48DuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ROhWB7fxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func duVvkngqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2G0YX45QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xZwQkHdzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OutjfVcaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func suq6iBDkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sugIqbAuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hp2IROHDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lfeeweX6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Irb67oiEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IPG95LumWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vrpLdn3PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nReqNWEIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s970meyAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LesoV11nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ykHMjYN2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d4ZxFsaRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VxsqI68qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4M1Vl8N5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hN5ZQH3eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VMI0jDfhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VXP9iyzaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1YZwDiIkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l3rMGpCEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B1vl71RVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func llIzw0WSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q86Ip8H3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s1OIbEyUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DyvFaWXZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qM2PcEt6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fKDqUvRbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1tSlzua7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2XIdNP2KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l9AVxwrUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oEg3RldtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 32mUW1E0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bw2jlLBsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SGrEvK0uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iAa8sZhcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vnOIYd0OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VpnCcCE6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func avgH4t9tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cOE1KIlzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 44Uj6WfZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2WafpQWuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ad7VC8pgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AWm12OqMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tFwt9vnXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n2HsnzuCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zTgoyEcJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KE7niR7hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1BZ0AulzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e3VTV5uJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JLAXmyLlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EMAIQIbHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OYGKogxZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y0PX8btSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nwxgvcQIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TFVmVvgqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c5C2xjZCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cqsJpAjsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Iq4AViDeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func taNLHmHFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NavD1JXGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5DrOXVTwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iDuSIALmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dOYq3KxIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pXuTAsBtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jAmCh7VbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lCEMkvbNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LxRtWKhVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f0N3Z32VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q75jPDNtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B00bfJlRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S5Xetd4jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ODVuSwvkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LdIq5pMxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Np6kyGW7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WBvjmVYaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d3fOhXGGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tvLQHlCbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wv84VRlSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1z8iaeGHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RS2N6SZKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 26IYXQxIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O9eSJ9FOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eyU0EPpeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lwjvO53tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3L2EHcxsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A1BbGu8NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bJc7oPbuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4MkZAOzlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6mXGpGE4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KUfY1741Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1vLtgqDsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EqTz3j7CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EjWXlxQJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Dtpr8djbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TmVTUh42Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4Ve2lmB9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HdPRcuhrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IlyWTfzZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4ATIkUrjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eaMjK71RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wjlwFqmbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eNuKRQgNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fGOSO72GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tPQ2tsvHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GAkN07xIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7TpF5jfQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ti6QYLE4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 07EUHWyAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func afY1VMzfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HVAN3CCZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OtBmpYYGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SPbdfUohWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PX5rfiRNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0Yg3zrwvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8118dsw4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MLOubFCAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rO6bAoXsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1MaRpU00Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ejo2hemzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6jq5bYLfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3OEbgum7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1hZSMo8oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0bzxfGLkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XHgE7DqKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P611EEb3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z85fl293Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ECXrp31NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GwxRNmaqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZpdXh8AHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Vhc4dEH1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jU0YqG0vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8VznDd9iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GMJNeyuWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RajPaL4KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Vr1GzulhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aREUJG1QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fLrZMvazWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 78UwUm7dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J6aVBsBFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CnZEmYBLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 69CTuONwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ys9vTR9dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func THHs1yBiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HNYi5YsvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5Mnr7UaXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AznM1wH2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CMp3r0vtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uPDYPU8vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VpCwmAOdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MLuHaKT7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ri1X1ZQZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JykK9T9NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UH5Td2r1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JeorVqd4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L4MjnHCJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KpT9arGjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6eYLxd9tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uYCXk3LIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uc4tGRsDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wPk0fersWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EbVPH9LeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Znm7ExEQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func stdzbytBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KjsULc70Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LzRw3z5NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func onwGgYboWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kbRnNVl1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7pFo4z9eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KbKYUSXLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0tk4oowNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4xR3d67EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rlQseZslWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4rn7FxupWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bgFyaWEgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CmcxfqiBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xcpRHNqTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PP3qGHH6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MpNeRoNXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZcWrDmwgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 45gAAFruWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QL3EUbyCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p1BNv4e3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ycEnZmUKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pXAIQOwDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kGhgRAkiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RIMygKlCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FbTDDuzVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1SMBRwDUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jD5GpAkgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cdQZMtZZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FoPbnm4hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ybxxGgnZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5uSSLSMAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4bru5MUTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hIjKHOPVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hYaeNhDdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7JKvPZUYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MdwaQK0WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uDZsSfvVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mTtHYUIXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Iqk8K09AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KHAHQNGWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PquBOpDdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ATrwznKhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b5blZDfgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wk4Uc4v1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wPpyv4C6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pqRYRLsMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jy62ays0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ILSDaFmjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xvFfYP6DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D7J2HYa3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Pzqv7KVSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func urUnhMkUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e3B80fIDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BtCi7F99Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ug79kR2zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xi5ho0saWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FCLuGdfxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1yedD318Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y49VqdJlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pfdAznb5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PBuJle5SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KyzioDU6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6o7ZS4HGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EFuxsBkmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QEg0t1riWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m6Qt1RKAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func crxfFGwBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uMdcLyZdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fuQFqtOZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kfDTTJyFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rulCGjOyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wRMS3hHQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BRGPWe0pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aOTZToDwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hiPhWkAPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 60Di6l0EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GBfLAd4XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fhBGJmpRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AjEchJz8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QhjhC3SEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y6x2PF2cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8aF6kjNxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pjnQkwhlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QekREnJkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Vu1NUyGGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yROplrMOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MmI71kNsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rXDHnOCRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func coXkfRA8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nZsKzVrGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8mqWNz6KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2cEwyytUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AuIuruJDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kYAh72trWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6KNtFRilWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QgZs8owFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lvWxwyC8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5oIbLbcuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tCJXiuWGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 49zxbq44Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M54xrYzrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zHEjv4l9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 91tUyxsqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CqeTa9RFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zixJZXMeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i5zLxdSAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9SL1m7mtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tkHnpldfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zACGvVyoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4zp4q6xbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XEDXAjG2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ySvH5zuDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gff8HhfNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iZKMhWtrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VJvx6ZHFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9oVxaxuLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cJAVdqMKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FiW1Ze0MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4a2b7vcJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iS517bwdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 65Zj6zS6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v08JSeWiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func trNRysAbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B3yjwY82Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tc1SywuTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sI1UwRB6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PH3q1mM1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2qdxhRoRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OiJkBvXJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HtQglEQBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w5Y3WuiRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KEcKW3pcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BT6REfI8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W1af2bWQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vyjnjnMOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Fy9lYVQVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qVzS3OeiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7inQf5ESWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UG59iMEbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bAusk5byWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l3JWpO5tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SCp74SBUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 50C54vOqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ih3BdyBiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ejvzpY5gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fv57s9nRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9lGJ72yLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VF3abD7BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AqLHeQHgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kiZotwHeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0PQhm7HXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0MAjVXmUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mxuQKAcGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xw1bYCkgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VR4TN3odWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PRharoN5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HblYnbSNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5wZzNtZKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GIaKU5QDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OzQ9Ryc2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DAMfiB7dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gVy3r75rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PRCOig1pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kaqwXXvLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vXqCkvVKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0fIUgUkdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ow5BA5H0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sGt0FKanWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func akulSkUJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HrnNu4djWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VzOahaBPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7tf4A13JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hrxgvzR7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GdcWyoabWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9EcNabv0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iwYdfUNbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lHSgcM48Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JdbFBlwpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xzhHxtbVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func awxwtG9OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qO8qbihIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7TJocri1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 99P15PN6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ffjUsiQPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8tlvvwdRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sXGzV3RiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iayYIPaIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VKFNJDdeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oEii6X0JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X0QCV1gTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K1Vl9ezVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JP9cuyiOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m5pnn1goWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GIvs7eMLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3oFQVoTEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H8CTHxvVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RCG4rucqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zziQ1RphWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mZPM1PYgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qoLCXvMRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GQvB3spiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8DxaeGzPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 72QzrazrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3sjPUAcYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pvkOdMRwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y3WhUwKgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bLtgUMosWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MsrOdFySWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KTKntESwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func llye8Np9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 31WEoa75Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HpUT1rMOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AfU907sCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func STuCRPijWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fX3QSG8gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yeOGqIfbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rENTTdAuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3gLR4WoMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E8lTgR5gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func svYj0LalWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DvJhyLg8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8iZAanuvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gafTMN6PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func azbvkPCCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v9MAzcBhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IaNkE9T7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M4WTzB3uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1iY6sw5nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ElLvWIhDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q55TwVxsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QhI9M2A5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ng6Urs4XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9Nl5JVxCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GfH07vjVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hW13cPTbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qeWWsEYkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3naTUhWQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a06IFrnsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oSvmiUHvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wNSgFodzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ain0JZyNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z7lDWta5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GRLF33IyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LISYepKNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oFE0KX5CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dBQpc5F5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BoDZYP4tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vgRd70yPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HJvKEvCvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NJhwgUZ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NT4TYSLkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rxBnivu6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V3lMdvn8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FwrhPGfoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aS77RyicWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R89Y6xJfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X0Kcto0dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XqXGw4Y5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S1ZW8mMVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lJgUKYiGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oxznTsFRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GdP6rzDzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VG4KcXFXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NPphfvV1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Qm524MVuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OPDxpTOqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zQ1CKNtkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LO3XYptnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZKQoLCXqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BmP74FecWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OZK3GaiKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cmfNRgRfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YwaayYrNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7tw1mrSEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eU5MClf2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bywZw3bUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9lgFgyFDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sRuCgNvYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func keT3l45XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rY77oS37Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mdNz245UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BScOW8E2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0mDA53KIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mF0e9MDVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tWt8r39sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HhHrm8jeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XAwzcTH6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func typMfGQBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ucMHICGfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4BLFGFwIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mPIaqh2NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4mdJrfE3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9US9DH5jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZMXHM6EaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wmibCgzmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7bYpdkbwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7QAW0pEPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yXbwjSYzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6eiIRNRWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XjYYXQ9LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CP5rQoTbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o3jL7JfzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2MKL6CH1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fXRUf7K9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KlJrCgdGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DQ6PMIMJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GEHNnqddWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d67MOKu5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dpmfmpdTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JhDOVjbAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lZp1ndJEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V2mSUx2eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FVvGXlUcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QWmmXOrIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ONhH2P9EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AJ79XRrYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gUdMrz9HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DHyOREetWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pfazlLTsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WqY9zgnlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jHi6LeF8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ce950bHtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HpHzN5rUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mb0IiIXnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oDrRCbzYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8RosMIj6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jR7f5mmdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZPDIAM5oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w4DNsejmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Xdyo12PvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e1repHprWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wfEfE2OGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4sxxFV46Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tjIGG4WcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mtksyXmXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X3qlLEksWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q3ujM2EyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WkOIG0xzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LTpktiESWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5z0L9V6zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1JYMB7mvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l0azsuj5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MNjEVdL7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func unnm7tOBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mxgRR981Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9UoannwuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cSsxhat9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y36PapPfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1Q8ZSWbDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kdA20zvOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func adh5G4dfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W2Z8hif8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wYZJgk01Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EomSFQo3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BRTSbfjoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oYFv14EUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6ArpeueqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 170NP8wGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uP7NrICWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FZVVOup7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pz1RVahXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MSkjor31Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8wyC6QlOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hB2nc2euWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xQq4fDkJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MGXuGJH3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZIsfMrEoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XPslB7ThWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BTYLwtGAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wjlcpKNCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4IQ0T4bFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SjVTM33TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WJRA3wjkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C6BBS97TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u7czkr85Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NwVg9oLJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1kUzlAl2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x8PFmhWlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SfpPZ3duWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pxKojzSmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2faGZiruWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DM0DfjQyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kkmVKUTPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YFerBDzVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YD5Ou71eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DC9GIBfYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KSQ7k0Y1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rP1kfACAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KL7G1l1kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GcrW5Eq7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OoR7OtAqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ygr5jSC6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HQ3JMJR3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2IqZIN5WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w2aTWEdiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UoksbxXEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8KMvr82WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6Qp73bAjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uO3r1uQuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AlsC9NseWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0F41ioY2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LI4t3xF4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HKJONDo6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 37pBTuOUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aSfOMGmbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9FvN1US1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wwO4mc6iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oxb1bsIdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xH3tUxPwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M8AttrqZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func unEJOioLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u7mJ79eqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Jw7kgNVaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7EWDLHX3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zq5G8HOnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PigmUHIxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jp16fSM9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dku7Q4esWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DmbdyjKkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YSVQAbozWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hy4wulb9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rot5agVOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tP39szpyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CUEZwivZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aSW86POlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wDd38wzOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QljSCwlyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YT3nh59WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ygphpmO2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LhMz3zf0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a8Fy61LLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ue4PXvdZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hXLkCm3tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E4DLDMhbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func drZ7jp1jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OCj1Rw5wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s9PvxpcBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QytaQrfWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2M54YoktWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 95IkdJa1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9bUkt8QUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WnxGn3bOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func duHdnDH8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Tt9hrX8yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oOP4x7d0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sIP9qcavWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jasnURmCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5AP2ChSZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VmOFNw9FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6HkC7UbyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EujovLi5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PxgEALITWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gOfYyBBCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HVMB9UpyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YyC5zwmSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Rb7Z9IKPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2w6LOwu7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BJBuLhsdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nWzk6xvpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WyNUk4VOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2yD9vIjjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RoIxKxa6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EUa7vosAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ffJG9LktWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ApAVfF4vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5zCQIKPDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func thokv4yjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u00geD4SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BJ6FPyVJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dY0K33X2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nncjy95NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DQOBEPGZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VCwwgJ22Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 84k0wwwWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4ZKRrwrtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x2CoVNXkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0KKFj6YVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CtgGK7iyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s9cwKIFGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6NZZmLxNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tEriXjOOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gZxig7r2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NDDncBfxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zDjabeyOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MPRUhKW9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H15XQnArWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iluKeocAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GEwlMNzXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mgq1hVvtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0CwdybViWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4qa8gUL9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vxStDTGDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Dxhw1p5hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y0zcz9ZcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pzGo2xcgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7mBd5sqNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l9H4CZv2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func No3SsnIhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PLZYC7gWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5TXRSe2tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b9U1mbnzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QP7VlaPZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2gEGPMk2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qFwdbHVTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 09LHezNsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aHHW56uOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yDOXGcbUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kPjYtPJ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func InDQtiEDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func coqwiFMxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3gMlutsPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QiPzhBRAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hb52jnW2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HEbE6hQcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AmzonSZ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KjieigcPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1BSyWFv5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2k1Kt05xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Hbxkf539Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kuukg2jYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PdblKq1TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dS7a1vrdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7UY6Vd91Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8j7nlmUvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4jtyZV7WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HpAzkZbAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qtzhzWs0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b7Bly9JQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hy2ixvbAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wQqzUqEUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 485J8EOwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U1BdtQYzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Pdfm11LgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nWQEVqcPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 96JbhdlZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8Q1oVXAxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ytHcgAUJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5VcJLqPJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KWuTMj8bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rIqnU1krWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pPCDzpUNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MAiOaJHTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WSwlaGbeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XUF8y9rsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q1eRh4p7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func spWiZ2CvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jwJqPZBwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 52wd9tvPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5f5pGm1GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H7TbmkEwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OxyBsGA0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c11jy9Z3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QO9MVB5kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ql7p4dXDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QIH4FWfFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AFaWOuZDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Rz1Cz1bbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ulinr18RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y0U5CC4OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6ZyW56u2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8GW95ZsbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 223j1Jf9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GjiOHv1dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jpr0PaCkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ATOI6caWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yt2V15JpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WVZNKZZWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ugaX4PtvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5a03EMBwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NRQ8EQlYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4lv9y1AtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hB2qS1fSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ShPjTTTlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2S2lumvpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IksYU9u2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0uEhLsKmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u5dujiudWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HqvpTM6sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s76yrIctWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func znSRZo1BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j3c46T2sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 35vSOMUcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 87zK5JuOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pdzLJx5ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9uqbo1oNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QAqtBpPoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FIZmDB1rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qRU5C7xYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2SsGRnStWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cH1M5hmGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bJOZ6QDgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sNQiUmgqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RKtNnC8mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j8MAN8OxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W1gaAka5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9E5GHc75Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GlS6irSyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xw66S6niWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hcsGrZaDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7unpqwqSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X9rCtUP8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FeFaCczUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Iy5VnX9RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func csJURnDcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1BeXMfSGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lw28iVXNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CN4olkUmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GGWoVhIEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4yxuNEKoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cS30sHMMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pE3p8BOXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IEQ9l51qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func noOqrSfLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XAOkp5j0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a5VzgYr6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 86VA12rBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Pa28s0VlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s7y5fq8PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iAB7Kmt8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func leJoU8hiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cvcTRHQuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FzXiEecVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SYOc5uzlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zBW5DRPaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cvKfDMFeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vLUSUyCzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v0th3WoIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z9EkwOG3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8Ri2stEEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q5FgucoyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sqMDf7RhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yDVN0FCRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MVUAO9LWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gmbSm19pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func POu3MDoJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yLEJS3sUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0hVYpcqWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func crODKjrqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NqihQPBxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2GZ4TV4GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0zcBzmDZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gu8yG5sbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QvntfBybWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rrlPxIVuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func idROvkbwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ybFC4S1GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AyLTIRMXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TQhso6fUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qGEBYmiLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5PuTks1UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BNJEHHXGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m7NtY0kvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bRNZsThjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qxw6NZXLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VtFDvm8TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RboyF8CtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wpUHzuBLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CBKMgEZIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8W36T0vbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rY8PRbYeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mQBofnYAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r0VKN7w0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aApkxah3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f88dnp2UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ONhE52OhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xluvpEDyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H6ORYCcLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 16KTrn2ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LBm4nKJ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6orlhMQVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e5MKELCeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gamaqfbgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YYv9vwyRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JoVXySk3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K1XGC5yPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0FZp3xJDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qYVZQlR9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xx4s9ruKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KisINGgWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qYgto19GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M0cZ5W8rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ab60tpqaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5RETRdydWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IaSNpUXJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dp4mf1DfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AHv4ywvqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ixa1TCEZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X6XlxRJiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ncm7U9D2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mN8E1961Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dI3UwPFeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bh3c6tR4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jWGATMJmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a6F2jTmWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jP8w0J6RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VPGI2JsHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VopQNApmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sxGbvmqVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aqV6Uf1FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func heM5f2jCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eCS5txq7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DOc3inh4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wte6THeJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mQFOSzYsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4PkDIPhZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W5pfPNsDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oqXPvj6IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KdCk3VnAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vkcnPM8vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RwKNJQHwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OLK4iYjIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PQU6B0cKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lNV6zPkhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Dw5My2EMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1dkrIiLlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 205YT6ZhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mUVy6nFmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7A9CiRVHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HWYAt4HXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LMYvMaS1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w39EFy8RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hLCYgwGiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CNSzwUhUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XL3qy0ylWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R2d2mrbEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zkpiRPcYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mzLn1bjSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wO4PZAEAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vsEmFwZsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AlvYaFNNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Teb7VDdEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9v0wwNfbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RruRRm9MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T8BxxJEuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FDVaXjWdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hJvEqeGFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xstyXokEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func enk7dSKLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wvNwXvgPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dniCl3eCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TitW8ObWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kfzZgjkpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vbsZfRAgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O1BE3869Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1Y6DksjhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Iu90Bd0KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6lAPpWrwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bfJ2a9hsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DwME5RPVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iq8k2i7ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ITcWg4crWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func heOOhFW8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ttZ0AKouWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hYDEbS03Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OQzzEpqoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YhGVKyQIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xWt8jLJYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d2id4asgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AAL0UzBHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EhmVwyrvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eXCewDmlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5RQW6CbpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m0S7HvPBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Hg9ER4c6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1Geds5B6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sSdWigtwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vGFZZyRbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qqiendm9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d8tQP7CNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CK3i7nc6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PzQj8iweWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kG4h6a2kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bTFFaaxwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 54Yiq0WaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FmqnrxlJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wnB2T1GoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yp9PATGcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I40ww4lZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bAV5BpKrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xBKu0FlFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n0167rMkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q6oRsTrAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hXK9rz7EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EzvHrkLUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QlBUl3FoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AqcPAGajWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kfOiYM04Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MTVZG8MLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Vi8X3DoGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2bbczUmBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F3TzC1GDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func waRN5BVfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xQAX3nTQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aGzdCrDvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RHfLvmK7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wupNcfkVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kcBMtZBuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cVlcxyGlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3ezSw8XuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gPXvxrGvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wY097ShnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d4Ow5PouWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bh4UdKnmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h6CVfSqOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 61e1mW2wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mFFPX3EVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T7I4AplXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9bKKugCyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8uNzUwc2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ztSZ3iTIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ta2tkhs6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HOtBQDeOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yzyBjQlQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xioH7T35Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e0524Va9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aWmZY0jgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mEw1zamOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N45goL6gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3RYjRJNNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PXY1JJGnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1O4Y5FNHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g3jCKFNnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MqCzUpmPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RwrIPU8FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5NBUP02GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gvmyQ3PwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dtrP15k5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sndI7gIlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Tiuj2LewWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7E2aeqXfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xQT1VsSOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EJIpc81qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QMFL6vRCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p9IBlYFUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vCZmd8UJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ngDMqD2pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bRUwHEwWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cwSAkNlVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pvED3QQ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qHucYtziWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xuQrIRTSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jHZLyn5OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5oLnO6FyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pRMQXJgcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M8uT5Id5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 75Rx731cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func osZImUUUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xW9zcvLOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OtnG3nu1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fNC27YgXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gmi41kqYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kWsralHqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BlbLRcjnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FkgpmkZgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LitLz1O4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7meCa60tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Mln4pQCWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M1IM2j1qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func let78i7EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z107dmSVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ypLDCjCfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eqAmgjWOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yf7B8lJXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aN6PmNNCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pZDGo9fUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9vMe3uD4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s1J7GVLcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QwxZj5V1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SOLli1DIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7Dhb0pHoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5THzI86OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uQrJe0XjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uE7r63tuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hsuXZPcZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UWeOZ2I7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZEtMwWJ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4TFR4JqgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h42kEdRQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T7rCaTsnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0NiDrYoKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ljq2FE0fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MJFRnMVJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Iczm0cmFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OGzwuTZfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ChkEf5NvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4EXmQFzmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Oy23RIwkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xHpxldscWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4MvZFGFGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VpWKvbzoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KX1953hNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QI13fRu3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UXzzxZ7JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hWpz3RsRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pqk1J9HxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n2txSsRqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P6Yzod5NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sStKWkvDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2Ecq2iA5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uw0uebAAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OSSsVLnNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tmAZ44uYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UI14388jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wlc5399OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ElK9RDATWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TXjVEgIkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5BKSeZAOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nfHIv15qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ggJIuMs9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qYJ88hBQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LbNBWOD0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5LwWQ8PAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BdRvv3AKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pQQJczNfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jT6uYnA3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4dKMgSiSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cwL2DSPlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j35P7XozWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func drys96jCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func siu6icGOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zvVnBzRlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bYA9YWCrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZwDdo5PMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uSMTnsuJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OLqrWnwJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bDZPjE5sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qQ6TBfHdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uYAYfszVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FWz8IHQGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3leZSHL1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D02UKxnaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q0MulPUFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uNzsXP49Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aJs823eOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3TGii5GjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xCTyHzszWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zm0ujbpvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Vu9Jaj9lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nAyVBlrQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Iny8K3V5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7WrexD8tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Fg8y1wiKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UDr7yku0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QKzSCKB1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9IO7i2brWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KLJ6kaIpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5WUvr84JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Mra7j9TdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jkf1oNsEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lmY3BFKXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V9FNaEZ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fVjjmHthWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4Gx1MoZBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qzscWIJxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ufrbBaUmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x6HT5KEHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GZRlBEsRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S55szrgHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RnLELHfAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QyZhynnBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4km294ReWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n6b71WXjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kCORaxirWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IeP9fNjmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JF85X3RIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rk3qOoXIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RVuZqhHGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TT9aNlXDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uxrmsXmbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UJH0iv3eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JrEfDcOnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W8Rqq9t4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tz0oSqYCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qKcgfiQCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KsoCaB2EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 81MFENSaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CvmYzUXVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0QSjrOFeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QJnanREzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hwISNh0uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2LAtHPf0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RlmTaHs6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PV0PpySRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WWVz9jhyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3idc2HRJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EHSRgvHVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gAet4NW6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FjCTqdjlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y9inCh3fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WntT3jHHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jJKVScZLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DdBHCtPLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FtThPBYGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hwas7Nu6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XvrstixDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J080bHhCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hr9iYmL5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2vbQyHiqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n8FgH4tWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sF8vF7MmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 83yQbdx8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t8SAoCwdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dGH9qiifWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0760qQX0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CDl5f13FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5kZhfIAtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fPOu3n5uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0g30TC27Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pk2kHqHpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uw3w2oXCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Vyh9571EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SfbvQ9KzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pFXaBLP6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Fus3RHaaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tPegmCvxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DnTXDF7kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4kLefAESWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KfauVUTXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i98LjFFSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LbHTQji6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Xsf2VmBrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nTZYskBdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F4AJgGcIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func szXxrwkCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fPTFuXnOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6IfAghYGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nYmIBBbjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ho9LBX43Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func omsO2kVaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func knQq4sd4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6CQFV4XLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aG4WEPhuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i6cJVmhqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QqeDDyotWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ps7AlItrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wR1ExS21Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m5uurbmRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Nx9mDBCCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func glU8aHeiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VPIOSCSpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Xjc2dIyLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LRx9QPUZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RfPn5xt7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0Sto2al6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DAY7vGcjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8mbF3LA8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z9awNtHZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u64pLfB1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HFGbzxBrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4d1KiYcfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qs1nfdIMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fRQvTcPSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G6l6b1KLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F5d917AlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UTCTCsZ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qt0yu8TzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0PgQlZORWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z6xBujiMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JJeZfYpOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d9EPH3vRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wlJOWTfyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DrzpeVU9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HvcOyNfuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7LrsNX1KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iAQKqJqvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZTlC3CdCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7epwiZtgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6BQsCgwVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1PdTBenGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SfHpOtnWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r9VZAHSjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func herPDgFuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cZFJprNKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jDQocLSuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ii9kf6RmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FT41deNaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vv1b3cPBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 30uTVTUmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ncIKqrs4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lctk8NhxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func erT7NirOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nIroKELUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iLDCChVpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OmF1vbIxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m0rhBD4XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aGV5MF88Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1rxjH4DiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FWUFIKySWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TMdFj7fYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FEJJRVJIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P750NGEnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 55IQLuagWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Byv8GuWKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2R3xE79BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sO7Ak8vsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZAfM1mnWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U4OSHUN3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4i4aJLg4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UQZcNtCkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 22lszuNrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a7Q81tayWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UrzhJ9E4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1FsK2WEvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EPvLzkFWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8vWz28FXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func alcPh78bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UlcEzh9qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RiZkgNQgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZLg3RXIfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ryApdhPEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zVy6KA5lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kCrZJP92Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QEeJS35NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7vfQUj5HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c9WxVyWoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8ACabjiqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func votlSu1vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8ZHBdMdoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LH4GHDH9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lI15wbWBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n6Pm2lD3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zicVKkBWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qR12kVCwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ds27CuvOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Li01nKgVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Qnifql6MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QVwxHbHKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EtQ3FDJNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fCDZXCQNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LkDW12lcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vU0b2JtFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IyNyN2RwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZfBL1p8LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nPslEKokWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Nh9OChxuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IC8eI4osWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EfqHsT5zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7uUicQgAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1PFa6AXPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eL7K8h8iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xrIZImYwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w4LVjqPYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2nLwY7CCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sqkyGOIcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R48NyJttWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hAYICeQZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r0AOIlCLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mHEQEYHXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RwG70zq6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IunWkEAhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0fBm2yP6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d5K9t4h7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fu2g4l3SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hu83eIlIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NfO6E5KJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func egwDv7vhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cXbcmIwHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TpsWkemIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Yj6n4kH6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Pdt0LTtiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UAdtHAI1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ads8W5uiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KBOiJrmWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2uTXSlJDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nWkSZaNMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Pxwwk1mKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 12J7EgJOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o7vJmAmsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OgcOR5juWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wdyb9V5HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fJmyqOWSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KovCPABEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func itUDJEIBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8hj7sBLUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pdMBMJcgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X3HCQT31Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lAtuIPKXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yYPpjmnfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mhxpzMWVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f8WiFyp8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9jW91qBdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GgIUmtuRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func emzsr6rZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XFOPpp6DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y95yqud2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lT6idiYNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VZqTjDrSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HKE8ndjkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zmlFqKVnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NVHhFy7bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9cxTu7iyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6brhDpIzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vxRTGJGzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NJVFz34CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7yEt0b2hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aWpqRhNHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JGhz979vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func arkWcNRmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n383H898Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CmiEZcWCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HAUsMjSMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k2nRHRxfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ozmJ3wWeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NHlCF3qgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CPqb8RZhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1K3rPsCpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0IXTxV59Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vPci5q2sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UbXofteKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yyZIQdR4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4ORgO2P4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XTQ2rMZ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func itIXLHocWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iAdOpdvXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mmeQ72CxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6SfrHZddWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k3IP5nSxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EHCAijFXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B9z3TaLnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9tOZmFhrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1VeDHU2bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BzFAoAftWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MNFG4USyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a8Kpnx6eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ENVJHHcDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ssY7UKWMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YsIxQtVnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 235ND7o0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KhTEagstWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zyw0SRg6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 17p9gao2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zeep6HznWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e2KQY44fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tT7JpCvpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EsmBQibXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ofZ8jXXnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xcpgcxrTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ozCsHVlAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vLU1Z3VvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5PQTGmXAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wOMKiPtBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8Ejzoj3xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tINqWctRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oVhlCghkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Fk7TCotyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vXVaEMr7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CUrTmnvKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HZRbmBi3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qxeePbAPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8g6Y2i1kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N7F8kootWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FiNCC6cdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bCoBgbImWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JMkjGV70Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OiI4rlkTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DDA6EuOUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8psKlVblWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yxIqhklPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DoO4OExTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qrZ0OIt5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0gLFfsERWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func br4qZajLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b4fYzZW1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ad0XGXFhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y89JlWpHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mDVr3mYXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cJtep7diWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cCC54JRQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aGbCnBXcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Bu4OKPGtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6qICR7rKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oPC4b9N1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z9FNzd42Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D3bPn2dHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func od58NFx8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func USkKD85rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O0TH3WmXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pwMtzc1ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IYPd4rvTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yGGI3ErXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XsSgQyEcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oxXOho40Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wCJLhdB2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nsbNpUyOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gOn10rjyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UVLrqLJtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VIkkCw3NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MsOIBoZWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ld3VHKMvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dVzspRBZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J1b3KehBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4ZwumkhiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6raxvtjSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p1SqqbgcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4Z3eX3U0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 03HKtFcsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N4Eq0aktWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TH2tSWV4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MiTICynXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I5FrRrPaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wfkU5S8fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NQHSrGiPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func igvPMoJgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SmjE088gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CmnM4TQtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cjUOivsFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z4zNLgBdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zMHP0M9AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dMGxjwzbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VJPl6WR1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9nFTxGa3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8kg7qU4OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TNUKWOu7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YCnndbXfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sqQJ0UU6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eYdTXQ9vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func THiD2tKcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wYQwR3PzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nkEpIwhgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KDYMEbZQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4nNHAdRMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zpg7T6BPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MItVhFEyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Eei6TdkRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IE5LRaekWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bh1nYxdlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2Qhj2amVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3XdNF6WVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SH6WXPrtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hu7MwNMYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 54J6WM4qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SEKnSQLaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vhXhUvyzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XSLLWxInWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mSKl8MxAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8ixmvJmlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2DgoYyy7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pClhldOpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QDPA9KeRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b5djt9QoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DLbY6N8KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PmRiiVKDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9s7YYfkMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QOBBcfNPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fWf1FFCzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lP74yZYbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZEcgTZ72Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9E25wCUzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UnwXKWAYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ktx6cgiCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vwV2X6SkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rvNGyw5DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J9yqJ6yxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yi9BsJztWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WTuj3iiLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func scEogd4WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YuOKXl3KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tNLLtfMwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OEkKxbgDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lJZ0nyuSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N2vFjAnPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q4jndMwGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9EODwVh4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QkJKwY0mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L1Ydc1VXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wloOuJhxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9C92LEPgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aSpSIanhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Nq1wVndiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GfQdX0M8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func orxuAKEPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 38jvK1ZPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6oP5PRBKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EVCNpuKyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HqSpt37yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jkPNKVKwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BDgRCK5WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RJzPansCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MrhJu5TMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U0LrjcF4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6zuW8nhcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vkMGNg8JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uNZF15e7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QqkVNprXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 81824Y3yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IsA6AFJNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kRRZPxH4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pkblPBzfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LCouN2wcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c5t2QToRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z9D5RKnPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o1fXAzGOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YXrSwiLHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WbBvqQb2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L3nf6nApWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wHoMaSjmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K9Wb1UXfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HGwD4YODWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Phy5RPBoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gLw7FumSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func crqF2VOuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3MSzYADvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s5xOUOuWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FdBwRVMtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3TPQpLFAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8BZzrRLCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JtQab6MDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7IR9zlxaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S4M2ecnNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I1yB9JzPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xRzivcnUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func om3IwJpxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s9IfGnUWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5WNz1S6pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gsrX6fG4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KhIekIMSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tuqwyad7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3FWyPB8cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CTAj6L0jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4IwBe8abWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X33EYISpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qFg1RMLCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K0WWJ9S6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aakMukyIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vcWfNNykWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OBhgxV9OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u22OdtE9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 265urLsqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rFrtp5QiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8o7LjCCPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dpyMqpeqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AaS0xjp9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zCpmjuXQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AQ9WxWUdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q8FTMBmsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VUpIXhVVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5x9LJYmpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FsSC17RgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MDk5KuvkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fP4RAisPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1W05OIBIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5JcYvBtHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WabSSWeEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5EYL5O5nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LvC4jM19Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GEn69UPyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aqZQUwzxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MQVl7beeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func agmlPaXqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Vqqi4RAZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xdLBLPAHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FQ0r04elWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9BEkLRCqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xCWPXLFLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LkgwSorvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VL8HpMlGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BkJUB1SzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L7t9aWzzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A6RC8B37Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 14VbbzlbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N8WIZ7nyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WUz8ou3VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func km9KbS6IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0Gz1Bra1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RoZZ0bKPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5mwt0LswWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iBrsWO08Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Zs0fF6KTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 85sfoYp5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BHsNUOmVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xCNZpvUKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c85LHBBsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a0P11yCpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zqFQVFkBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bd3g5HNbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RkubCJEDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uAKWDMHXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2k2hCaIUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JjX0TtkMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MEKQtDfaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ff9jXA5NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VPgJdFwsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F5kBovQmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FBNGh1t2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2gVokLtSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 61sOOM6SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GE4liJghWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f5cUusqBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oSRVF1RzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UJcjpOeoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kfdbVsUAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BZQKZ20CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wtWcL2VdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BBf5lFiFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZKIKWG8NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SayyoIBUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eGerlKH1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SYIiHcmCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yktiI3s2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kTwkIctTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KRymJNqeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qsf410xmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gpTGHEP1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q8eOBDyZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k8UgQlU8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hhzKhhPmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lfUkQCg4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0rWBnfNtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D61i6C7bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pTkns2CIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kprozBJCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v7P3LfLtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nk9NTvBYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r1pD7qHeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MdPeITC6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OwdNcejIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u60wB2WmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z1x0ahcBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mkoeH6LqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fP7AQpAYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PVpmUhL2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vzK866WGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p3DRuhZxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pIbveXjnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AXYYwazYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UdPwtFNcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fMRQw4N8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rYTb8KwJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4t5SQIMyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AU2pIkKYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DDDoHaOvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mQvFajtCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func om9FFEDWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M4zTwIm8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R1Se2BtkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XdGs2weJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hZe6nsNHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JhLfAJ52Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6OPJUpy8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L89ytqpEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PCjGKbEHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nPJD06XbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BSGpdT1qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WuOL07iPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wsHZKEhKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G15pwTeXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TJBtWM0SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FmkzlY7VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SD0mdjiGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MBt5NlCbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cYZTqzrvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2XPhPcaTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VMmwYGGYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QfasxWxFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UPtv0SEvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oieJHWZeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tuj6RJNwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tkJuY0a7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gvfJZxaFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MwMIbWGVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cXckI5xqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EncR9Gt8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hNELcriiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9HBkgjc6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sDC8uj9cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2Z0jFeqzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JbheQ8E3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VUpi4SbtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iJbcZtyFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lW7od68lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tjy9IOIqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CSPIqh1LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c2mkPVHhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y8Mtp9paWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wJ4kqEUYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xJWck10tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Cs2i21S8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2OCuts67Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l3IHBzcnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AAsGkSqVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6lIrPzTzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XoeLCOwCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VP53BqzxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XaJ1GXfCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tHiSpY6eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yTMXnFmJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8cWDNBABWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Vq4BzeZ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KCQvORReWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i2nInQKCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aq910ABNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NNeKlpYFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Oweg38NGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qqdNwjw5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wrPzpvwiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c0KgbUcPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ShuDUWlAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N2oIrt3xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xl7g4PnoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FFIg1H5uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OD2WJQDWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tn3YvMsWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d9ShShTtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func phZSMDdpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k1z1KQ8wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v2uIHs9SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0fGj3QHgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SpPuAXpzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OylOoaGHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JNHnCaQmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qNqVHVESWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8atuqfazWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PBLIROQAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QRwUi6BXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fijJ6PV1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jknAPjtBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bJKY0SJLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d7qZudDeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l7QkgPgWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wCRT7FSMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2a7A7KuMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vuYffwwIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 44l5fpeeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Glckh5HKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5M1BSQD4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T6Duh6MuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kjvuiakGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MtozKcOTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yPd12VcPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pnP6BLNiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func padhr5MaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DrUJ5MkFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ckhqtmzfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HX7ut406Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qkVKQb75Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R8w4HKBxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h90nf3PgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J4iPouWYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U623a5PNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xl4LdUexWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sVBoqawfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vmWbrWYPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vI9Lly3gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O3gImCj4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A2o8sQzSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OFf8kkm6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3YS0lV7mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cOUHw0ySWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ccill3FpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EEs7RYbdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OCGbqRpXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jlIJ3s7EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VtaNjS0rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0BNneOzuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OYevNl2AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IPgk1cx5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jeoBkj3AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Rsoqnx1HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jgugIlq0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G8bbsWs5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func taTnUDAAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ABaCkFuHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HMtVZPG9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VXQ3OiiPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u8ohGdDWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9YNOCHVMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m21twM6YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KYUaQPElWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2xcKzRXTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ngqZqJl1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hWyg8a5FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ctANtDa4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ujVMEWr9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bShEQrbxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U4J8u8lOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func enD6MGyTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wWlnLAhaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V9YFpqb4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fXwNY4FXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j26oghSpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QUPspMjiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ApvzxWYBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xSBHf8BgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j4AuKniaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kvSfTGmKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3hOrxWycWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oCHDGAdvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HMaE8PvyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V5zSmfGVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AerPx5wlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R6NY7DtRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LJsr38zVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ndJU4euTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UGZ54XmAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rNbs4XZ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ajbegS2CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wsph6lCtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QBBu2710Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0ZpAH8rHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7OqCLS1vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O250fxfBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JAl2gBpfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TzQDTNbqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oYdVdZAkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NLZZ6QtMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3nwCmQTUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p42rGwjRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PKfV0KymWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HfAw0OggWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VQdDGnfBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JWS3Wpq5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FxLdsdzQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CZY2KvpiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cT4TKta3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PLsKk0hXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CBTW7d3QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YvQJr6xhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lJ4VvObTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Bx7DobYmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NF43vFwDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EjZri0YgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HgjDkdNPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func woQXKSVIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QetJnT53Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K9AHNGD0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LQV3M9VuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CrgYIYSpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xveBjAOmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rOAUBYH7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fnLoZrrbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IPUfBfDqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 60t63sFLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Mrzd4yxfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Er37oj0GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2Fl2oPMiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uDEwPFUNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rgHiYTNjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kgDMpn5WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eDboUR00Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5YmbnO9fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WrkdParUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YRK0kJ2XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1Aog8nowWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xPjG7VqUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QHfvW6x6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W63QPkj9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7Gl1ORqBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xHpWOdKBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LcyqKl36Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7ceZ5iBoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qFS5qKg3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AcaDvtsqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gvOoozfqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MDw3leQeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0A53zPYeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E3018drBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Lr81HqXfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rlW9SmXwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 29E7FKX3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O8cWn4H0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func my1TDy5GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func taMKuL95Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1YULbRXpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XzPaptoJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ghWUdyhlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e4OWofFPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CBHopYWkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NRTPWARSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6DAHbEBpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func svgttqlFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WOcD5g0vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iOaYvyvAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P9ELASoOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ATWMc87bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func clGuv8K3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n6vCDC6kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SWB53p07Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kaV4gYpXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lP8f2jo0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YjPXC6i6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2JIHkl8yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZXl5nvpcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zWvHrOYaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cmiTGwXCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Zx7pRxm9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func turGkwrxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eEHS5cwrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func brMAE24nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AMJeJn3FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DFmey36AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lb5Q6G5MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GlFcWRfuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 45ptnYD6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mKL92OolWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x8XSoYdwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tTXrToXHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ONbKQScSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SCWt2upqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J9lGQUonWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func diHI3YmIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KS7QPVV2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L0WVKTg2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VItjk8gKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dwZ8E4OqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ICCDwo4pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a4Ga710nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gI9Ifg2kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FnextaEUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5Yio1UK9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fUrQcB1TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HwdPaQU6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pZynDDvwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4y4OPhWrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CrK41SjeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hvxMvTmLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H72SOb5zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fDjtUxY5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QqDj09UXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gzNvp3iLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GjtgIqW3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z4beyBytWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RHEXXKwVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qUMiStuOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ra3WaB3XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gcr060B5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VNbyutOJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7sH0SqrbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oX41i1iZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7G0nrsyBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lJqzOehmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vt0w30mnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TaSCf013Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zF4n0EC1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y4WVbWDsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vyHxDKMKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fSbtJMHXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lzUYB7pqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MUkDxWrMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cICcpbCBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gz0TfYo6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m64d3ypMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Qqw9ovfQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jJWv86AqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fl1ek7u5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yMEDlS12Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C9Kt9rJ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XToQFneUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func puYvkCGyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ksZRfJ0jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1NrcrhajWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func waDf0fs7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yoiWBjueWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xkJvFCgvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NhUYdCclWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1PhFv3LdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mIkfLFqcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func boCvfnGRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kY1oG6yIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v3cVAdJEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tHbjVsLIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ShwidEjKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T4HHDBhOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TsAeDig3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IasRBFqUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k1xnsIF8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7zaY0hZCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y2huR7XSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aNDiiZDyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qZIpFqJnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2hhaYwjSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ApKf5wscWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gz592HsgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RsNIQ5rGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3Y18DEDkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IXbiB4RSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vpGNODvzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T16YsvbTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BN8g53CfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V2A5NoRRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d5E0BdQbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 74sWlTCoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 63AuN5UKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YcEO9QkSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tXW0d7mhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NFlw7jiNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A785fo8mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w6YHr6mRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O4iSjBOSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AEZMhvs0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d4HM3SWrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NZLcShwZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FSm7TE7OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func shtoDnAvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func plLLg5VqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vNp1XT9oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E7b9MxEhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8TnVj1ysWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func keyAaqFJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 53kp0W1jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EzExRSDxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SMVMQQwkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YTSziEUUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func srPX5mivWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3BSLfiRLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T4aDEOjAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fS5DA5p8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1UlitVngWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xnkQUE9DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4YIRLTZiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 56lGPztKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eG4YVoaqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func INK53ytgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ndkw1pgAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lefSRsrxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6BqewLm7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D8r3z38kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jSugikqMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bZ8qWewEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func naKKLympWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2VhhlBocWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func clT3EHaGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qqKgOftsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jWJk0hh4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hMhyj25wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oOna95IrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8ilhNCesWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G7WHY8l8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qdI1r7J9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EKUznV3MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M6xeLmKGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yslvRYXXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f5mAOvyvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V9XzHGfeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cEfXVk08Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Dzrk6dQaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UKXCpyJkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AauaDSDdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N1xa7GCyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W4GNl0iHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pAHa3yLgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2yPVy9UGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mk8aKkOTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Tq4iFyFoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c6VgMHg7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KDsYJ0JkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u9G189m5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D5DirD7TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4KqmW9CwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SyNnOQkBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PQQSjObSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MdcjjyiLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Qd2qPYqCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hxvzdpGmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4LQQpA4JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 13GbXsMPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w6qGufWiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2kW3KgoXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vflGf37QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func guh2IK5UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7yEuYWUqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func umQw0xhfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o4i6P0qfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NTHQldfNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Rpi6ppbkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pxLC2m0NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0fP1hm8fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WjkjRWwCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GV5nBpFuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cZWAMCY0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Vy4gmzWEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X5ai3PBjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Inepi21BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nwMmxan7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m6qaoxVNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SGvwvN7HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bFcqtMYFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 62W5OC8vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ANgAYeSmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q9HYFeC3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ahbAklTfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JrlSjw9JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Pi6JoiwkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mjuwm7ryWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b1rUwpKpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ISGekV5UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kgRKIXeIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qKeTizkXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DeOClNYjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nnoLwTNkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4GCLGAf8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gq7jmwxFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IJUgdpWUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IjjFnmgqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d0CsBem3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vUbsgdSAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CB4NzoxXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7xcqwzrYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func glQYHJMDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sYJ30mE4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cVWCrAf0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E4D2WWJ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ek50UgFOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 57RI4OEGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xUNeGJX5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SqQbSXByWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rp073jKwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eui2CH5MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FpjsVYr0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uh7mgM6NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mMfwnObOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1Jkcgq7kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 02tl1qimWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V6lnMJx2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7CWefdQeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qU6VdQYeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dwO4g9d5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gRWA2X7zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cKXo1rxXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jGLOKSRvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oIqlBYopWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eCayQ3zIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lf7pbdI8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RDOXRiF3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f9gKL0mwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tbyjuMnwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AFaqvrEsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JN69bYD3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k7wEosgRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7tcGC45fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mvB3AAxDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HqZSsjTvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9wcFc8gUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func koIpQ1oTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UhCDyWauWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KHR2V4U6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qlyosqLIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KmvbqwMuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JvgWhmkOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6Pp3NXyVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gs6UapBMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SfNN0cAdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oo986pcjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BHcwLR5FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c49ySDvGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aF55BtJHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1kZKvD7AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hQYCy41QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kX0eqrONWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zpC30L7uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hjLutqp0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Zq7x3CszWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y58xtI9lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EaO813S2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jL3Rh2xmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Atbz7JkPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func to3OQAhoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HzBw3MfXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S0tnK2gkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LVxUec8sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OCyx1nZhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ov8tDWajWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bqHwTd00Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fM0GBfh4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BDf3itdNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KhyzQRqPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9EQ8riz9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func soAM72CaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oKD8ipkMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JkrFpo3rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wKsixsRVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U2TTcj5SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ASECjJHQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CtfGdL2cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dKV1Bzh0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PYg6DE57Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6IKZ6jJcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w8p39U2HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vCm3rnxsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fo7GKghvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AGBiF1DQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AG82g1mzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hjfTJg2tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5jGRcilmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xwWR1UiIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tbIwkWVWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TAtkbF3qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yj6fq7ZcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6QS9735JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1tXYWclKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PoZxgXxaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NsnHP6bLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GKXF4DPmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pGblhQysWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r1blBjKQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FulBBMP1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nIeXWS9ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tOWQfMM1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qIov1wM6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 891T9QN5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KvPYNRgoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QGXGHC4GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aCU0VmRwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func srm1FwQNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Oc9u6SUGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dBEC8oiNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func atGiF0RvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iYvt4KfKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e90UbuczWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KHcmwDKaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DcTptpUgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WeLcSJpkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3Ar2erKAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

