import cv2
import numpy as np
import onnxruntime as ort
import sys

"""
Using pre-trained onnx model to do inference of classification, 
reference: https://github.com/onnx/models#image_classification
"""
def main():    
    # get image path and model path from command arguments
    if len(sys.argv) < 3:
        sys.exit('please input args. 1st is image path, 2nd is model path')
    img_path = sys.argv[1]
    model_path = sys.argv[2]

    # image preprocessing
    # load image with format of BGR
    img = cv2.imread(img_path)
    # convert to RGB
    img = cv2.cvtColor(img, cv2.COLOR_BGR2RGB)
    # model requires 224 * 224, so input image better be square 
    img = cv2.resize(img, (224, 224))

    # some transform to fit model
    img = np.array(img).astype(np.float32)
    img = np.expand_dims(img, 0) 
    img = np.swapaxes(img, 1, 3)
    img = np.swapaxes(img, 2, 3)
    # print(img.shape)
    # The images have to be loaded in to a range of [0, 1] and then normalized using mean = [0.485, 0.456, 0.406] 
    # and std = [0.229, 0.224, 0.225]. The transformation should preferably happen at preprocessing.
    mean_vec = np.array([0.485, 0.456, 0.406])
    stddev_vec = np.array([0.229, 0.224, 0.225])
    norm_img_data = np.zeros(img.shape).astype('float32')
    for i in range(img.shape[0]):
         # for each pixel in each channel, divide the value by 255 to get value between [0, 1] and then normalize
        norm_img_data[i,:,:] = (img[i,:,:]/255 - mean_vec[i]) / stddev_vec[i]
    
    # load model
    ort_session = ort.InferenceSession(model_path)
    # name of input and output
    onnx_input_name = ort_session.get_inputs()[0].name
    onnx_outputs_names = ort_session.get_outputs()
    output_names = []
    for o in onnx_outputs_names:
        output_names.append(o.name)
    # do inference
    onnx_ret = ort_session.run(output_names, {onnx_input_name: norm_img_data}) 
    # onnx_ret is a list, res is numpy.ndarray with shape (1, 1000)
    res = onnx_ret[0] 
    idx = np.argmax(res)

    # get lin from 
    i = 0
    for line in open("imageNet1000.txt"):
        if (i == idx):
            print (line)
        i += 1
if __name__ =='__main__':
    main()