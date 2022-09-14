# Raytracing in a Weekend in Go

This is a Go port of the [Raytracing in a Weekend](https://raytracing.github.io/books/RayTracingInOneWeekend.html) book series by Peter Shirley. The port will be completed by following the book and the C++ code as closely as possible. Obviously, there are some differences in the implementation due to the different language and the fact that the book is written in C++. Nevertheless, the code should be easy to follow and understand. I chose to port the code to Go because I wanted to learn the language and I thought this would be a good way to do it.  
After completing the port, I wish to add some features to it, such as a GUI and a way to render images in parallel.

## TODOs:

-   [x] Port the C++ code to Go
-   [ ] Parallelize the rendering process
-   [ ] Make it customizable with command line flags
-   [ ] Add a GUI

## Prerequisites

-   Go 1.16 or higher

## Usage

-   Clone the repository
-   Run `make build` to run the program and the image will be saved as `image.ppm` in the root directory.

## License

This project is licensed under the terms of the MIT license.
