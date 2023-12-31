# Mouse Plunger

![Logo](build/icon.png)

Super-basic mouse control for the plunger in Visual Pinball X using vJoy.

## Setup

### Install and configure vJoy

Install vJoy: https://github.com/njz3/vJoy/releases

Use the Configure vJoy application to create a controller with a Z axis. Other axes and buttons are optional and won't be used by mouse-plunger.

### Install Mouse Plunger app

Download and extract the latest release: https://github.com/Billiam/mouse-plunger/releases

Copy `vJoyInterface.dll` from the vJoy install directory to the mouse-plunger directory

Run mouse-plunger.exe

You can use vJoy's Monitor application to verify that the Z axis is working correctly.

In Visual Pinball X, configure the controls to use the Z axis for the plunger (this is the default configuration).

## Usage

After installation, whenever the mouse button is pressed, and then the mouse pulled down backward, the vJoy Z axis will be updated.

Release the mouse button to release the plunger.

## Limitations

* mouse-plunger just reads the mouse position, so if your mouse is at the bottom of the screen when you click, you may not be able to pull back all the way
* mouse-plunger does not know if you are actively playing Visual Pinball, or using a menu, or indeed, if it's running at all, so many mouse actions will affect the Z Axis

