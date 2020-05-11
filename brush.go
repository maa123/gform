package gform

import (
    "github.com/gonutz/w32"
)

type Brush struct {
    hBrush   w32.HBRUSH
    logBrush w32.LOGBRUSH
}

func NewSolidColorBrush(color Color) *Brush {
    lb := w32.LOGBRUSH{LbStyle: w32.BS_SOLID, LbColor: w32.COLORREF(color)}
    hBrush := w32.CreateBrushIndirect(&lb)
    if hBrush == 0 {
        panic("Faild to create solid color brush")
    }

    return &Brush{hBrush, lb}
}

func NewNullBrush() *Brush {
    lb := w32.LOGBRUSH{LbStyle: w32.BS_NULL}
    hBrush := w32.CreateBrushIndirect(&lb)
    if hBrush == 0 {
        panic("Failed to create null brush")
    }

    return &Brush{hBrush, lb}
}

func (this *Brush) GetHBRUSH() w32.HBRUSH {
    return this.hBrush
}

func (this *Brush) GetLOGBRUSH() *w32.LOGBRUSH {
    return &this.logBrush
}

func (this *Brush) Dispose() {
    if this.hBrush != 0 {
        w32.DeleteObject(w32.HGDIOBJ(this.hBrush))
        this.hBrush = 0
    }
}
