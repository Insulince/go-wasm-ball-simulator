# Go Wasm Ball Simulator

A simple simulator of a ball bouncing around your screen that is affected by gravity, restitution, and friction. It can be dragged and flung around via mouse.

## GitHub Pages

This is the GitHub Pages branch of the project, which is used to host a live demo of the project at [https://insulince.github.io/go-wasm-ball-simulator/](https://insulince.github.io/go-wasm-ball-simulator/). Since GitHub doesn't allow anything to be run on its hosting, we cannot use the `server` from the `master` branch, nor run any other build tools, thus in this branch both the latest `out.wasm` and `wasm_exec.js` have to be checked in (which I absolutely hate, but whatever). If you are interested in how this project actually works, please switch over to `master`, instead of investigating here, as this branch does not set the project up the way it is inteded to work, it just hosts it in the minimum configuration to comply with GitHub pages.