# Setup References for Jack model
This guide is to set up references for animating Jack. It is not required to [render the animations](./render.md).

Since we do not encourage piracy, this repo does not include assets / sprite of Jack to be used as references.

To set up references, you'll need to prepare the reference to follow as this format:
```
SpriteReferences
|-- {animation_name}.{frame_count}
| |-- tile.{perspective_index*100 + frame_index}.png
|...
```

For `perspective_index`, current configuration is:
```
0 - BACK_RIGHT
1 - BACK
2 - BACK_LEFT
3 - RIGHT
4 - LEFT
5 - FRONT_RIGHT
6 - FRONT
7 - FRONT_LEFT
```

Then, you'll need to open `Jack.blend` with blender. *Make sure `SpriteReference lives in your current working directory*. Then, go to `Scripting` tab, open `setup-references.py` and run it. It should take some time and generate scenes as many as the animation given.
