import random
import sys
import cv2 as cv

def showInfo(width: int, height: int, fps: float, fourcc: str, suffix: str):
    ''' display video information '''
    print("resolution:", width, "*", height)
    print("FPS:", fps)
    print("fourcc:", fourcc)
    print("suffix:", suffix)

def chooseFourcc(suffix: str, default: str) -> str:
    ''' choose fourcc according to video file suffix '''
    ret = default
    if suffix == "mp4":
        ret = "mp4v"
    return ret

def setOutWidthHeight(originWidth: int, originHeight: int, operation: str):
    '''set output width and height according to operation'''
    width, height = originWidth, originHeight
    if operation == "resize":
        width, height = int(sys.argv[3]), int(sys.argv[4])
    elif operation == "rotate":
        width, height = originHeight, originWidth
    return width, height

def processByPerFrame(capture: cv.VideoCapture, writer: cv.VideoWriter, operation: str):
    '''Each frame is processed in the same way. there is no relationship between frames'''
    while(capture.isOpened()):
        # read one frame
        ret, frame = capture.read()
        if ret == False: 
            break
        # process one frame
        if operation == "flip":
            frame = cv.flip(frame, int(sys.argv[3]))
        elif operation == "rotate":
            frame = cv.rotate(frame, cv.ROTATE_90_CLOCKWISE)
        elif operation == "resize":
            frame = cv.resize(frame, (int(sys.argv[3]),int(sys.argv[4])))
        elif operation == "gray":
            frame = cv.cvtColor(frame, cv.COLOR_BGR2GRAY)
            frame = cv.cvtColor(frame, cv.COLOR_GRAY2BGR)
        # write output
        writer.write(frame)

def randomChangeColor(capture: cv.VideoCapture, writer: cv.VideoWriter, fps: float):
    ''' random change color in random duration, by channel level '''
    maxDuration = int(fps) * 2
    # some random parameters
    duration = random.randint(0, maxDuration)
    channel = random.randint(0, 2)
    value = random.randint(0, 255)
    count = 1
    while(capture.isOpened()):
        ret, frame = capture.read()
        if ret == False: 
            break

        if count > duration:
            # generate another pattern
            duration = random.randint(0, maxDuration)
            channel = random.randint(0, 2)
            value = random.randint(0, 255)
            # reset counter
            count = 1
        
        # process current frame by pattern 
        frame[:, :, channel] = value
        writer.write(frame)
        # update counter
        count += 1

def main():
    # usage 
    if len(sys.argv) < 3 :
        print('\nusage: python videoProcessor.py <video path> <operation> [options]\n'
        'available operations are:\n'
        '\tinfo\t-show info of the video\n'
        '\tflip\t-flip the video. A third parameter is required to indicate the flip axis:\n'
        '\t\t0 for vertical, 1 for horizontal\n'
        '\trotate\t-Rotate the video 90 degrees clockwise\n'
        '\tresize\t-resize the video. another two parameters are required:\n'
        '\t\tthe 3rd parameter to specific width, 4th to specific height\n'
        '\tgray\t-turn the color video gray\n'
        '\trandclr\t-random change color in random duration')
        sys.exit()
    
    # get arguments
    videoPath = sys.argv[1]
    operation = sys.argv[2]

    # get the video
    capture = cv.VideoCapture(videoPath)
    if not capture.isOpened():
        sys.exit("cannot open the video")

    # video info
    width = int(capture.get(cv.CAP_PROP_FRAME_WIDTH))
    height = int(capture.get(cv.CAP_PROP_FRAME_HEIGHT))
    fps = capture.get(cv.CAP_PROP_FPS)
    fourcc_raw = int(capture.get(cv.CAP_PROP_FOURCC))
    fourcc = "".join([chr((int(fourcc_raw) >> 8 * i) & 0xFF) for i in range(4)])
    videoSuffix = videoPath[videoPath.rfind('.') + 1:]

    # show video info ot not
    if operation == 'info':
        showInfo(width, height, fps, fourcc, videoSuffix)
        sys.exit()

    # output arguments
    outFourccStr = chooseFourcc(videoSuffix, fourcc)
    outFps = round(fps, 2)
    outWidth, outHeight = setOutWidthHeight(width, height, operation)
    
    # output writer
    writer = cv.VideoWriter('output.' + videoSuffix, cv.VideoWriter_fourcc(*outFourccStr), 
        outFps, (outWidth, outHeight))

    # process operation
    if operation == "flip" or operation == "rotate" or operation == "resize" or operation == "gray":
        processByPerFrame(capture, writer, operation)
    elif operation == "randclr":
        randomChangeColor(capture, writer, fps)
    
    # release resource
    capture.release()
    writer.release()
    cv.destroyAllWindows()

if __name__ =='__main__':
    main()
