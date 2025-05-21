# Hand Keypoint Detection using Deep Learning and OpenCV

**This repository contains code for [Hand Keypoint Detection using Deep Learning and OpenCV](https://learnopencv.com/hand-keypoint-detection-using-deep-learning-and-opencv/) blog post**.

[<img src="https://learnopencv.com/wp-content/uploads/2022/07/download-button-e1657285155454.png" alt="download" width="200">](https://www.dropbox.com/sh/5u5mf4l3gepe0m0/AABWoKDnuTdyABmKKXgFFR2-a?dl=1)

Please run getModels.sh from the command line to download the model in the correct folder.

### USAGE
Возьми `pose_iter_102000.caffemodel` [отсюда]("https://huggingface.co/camenduru/openpose/tree/5e17f6ad43ab415a0114537541a8d37d2503424f/models/hand")

#### Python
**For using it on single image :**
python handPoseImage.py

**For using on video :**
python handPoseVideo.py

#### C++
**From the command line :**
cmake .
make

**For using it on single image :**
./handPoseImage

**For using on video :**
``` 
g++ handPoseVideo.cpp -o handPoseVideo $(pkg-config --cflags --libs opencv4)
```
```
./handPoseVideo
```

# AI Courses by OpenCV

Want to become an expert in AI? [AI Courses by OpenCV](https://opencv.org/courses/) is a great place to start. 

<a href="https://opencv.org/courses/">
<p align="center"> 
<img src="https://learnopencv.com/wp-content/uploads/2023/01/AI-Courses-By-OpenCV-Github.png">
</p>
</a>
