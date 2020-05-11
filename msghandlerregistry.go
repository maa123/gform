package gform

import (
    "github.com/gonutz/w32"
)

func RegMsgHandler(controller Controller) {
    gControllerRegistry[controller.Handle()] = controller
}

func UnRegMsgHandler(hwnd w32.HWND) {
    delete(gControllerRegistry, hwnd)
}

func GetMsgHandler(hwnd w32.HWND) Controller {
    if controller, isExists := gControllerRegistry[hwnd]; isExists {
        return controller
    }

    return nil
}
