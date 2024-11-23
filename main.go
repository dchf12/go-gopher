package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework AppKit -framework Quartz
#include <AppKit/AppKit.h>
#include <QuartzCore/QuartzCore.h>

NSImage* resizeImage(NSImage* sourceImage, NSSize newSize) {
    NSImage *resizedImage = [[NSImage alloc] initWithSize:newSize];
    [resizedImage lockFocus];
    [[NSGraphicsContext currentContext] setImageInterpolation:NSImageInterpolationHigh];
    [sourceImage setSize:newSize];
    [sourceImage drawAtPoint:NSZeroPoint
                    fromRect:NSMakeRect(0, 0, newSize.width, newSize.height)
                   operation:NSCompositingOperationCopy
                    fraction:1.0];
    [resizedImage unlockFocus];
    return resizedImage;
}

void setCustomCursor(const char* imagePath) {
    @autoreleasepool {
        NSString *nsImagePath = [NSString stringWithUTF8String:imagePath];
        NSImage *cursorImage = [[NSImage alloc] initWithContentsOfFile:nsImagePath];
        if (!cursorImage) {
            NSLog(@"画像の読み込みに失敗しました: %s", imagePath);
            return;
        }

        // カーソル用に画像サイズをリサイズ (32x32)
        NSSize cursorSize = NSMakeSize(32, 32);
        NSImage *resizedImage = resizeImage(cursorImage, cursorSize);

        // ホットスポットの設定
        NSPoint hotspot = NSMakePoint(16, 16); // カーソル中心をクリック位置に設定
        NSCursor *customCursor = [[NSCursor alloc] initWithImage:resizedImage hotSpot:hotspot];
        [customCursor set];
    }
}
*/
import "C"
import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	// 入力されたパスを取得
	fmt.Print("変更するカーソル画像のパスを入力してください: ")
	var inputPath string
	fmt.Scanln(&inputPath)

	// 絶対パスに変換
	absPath, err := filepath.Abs(inputPath)
	if err != nil {
		fmt.Println("パスの解析に失敗しました:", err)
		return
	}

	// ファイルの存在確認
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		fmt.Println("指定された画像ファイルが存在しません:", absPath)
		return
	}

	// カーソルを変更
	fmt.Println("カーソルを変更中...")
	C.setCustomCursor(C.CString(absPath))

	// 一定時間カーソルを維持
	time.Sleep(10 * time.Second)

	// プログラム終了
	fmt.Println("プログラム終了。元のカーソルに戻ります。")
}
