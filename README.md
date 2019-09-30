# TilesetSlicer
Slices tileset into individual images.

Basically it transform this:
![Image with tiles.](https://raw.githubusercontent.com/adam-szymanski/TilesetSlicer/master/example/tiles_example.png "Example tileset")

By running this:
go run TilesetSlicer.go -file example/tiles_example.png -w 50 -h 50 -out example/tiles -out_filename tile_

Into this:
![Tile 0](https://raw.githubusercontent.com/adam-szymanski/TilesetSlicer/master/example/tiles/tile_000.png "Tile 0")
![Tile 1](https://raw.githubusercontent.com/adam-szymanski/TilesetSlicer/master/example/tiles/tile_001.png "Tile 1")
![Tile 2](https://raw.githubusercontent.com/adam-szymanski/TilesetSlicer/master/example/tiles/tile_002.png "Tile 2")
![Tile 3](https://raw.githubusercontent.com/adam-szymanski/TilesetSlicer/master/example/tiles/tile_003.png "Tile 3")
![Tile 4](https://raw.githubusercontent.com/adam-szymanski/TilesetSlicer/master/example/tiles/tile_004.png "Tile 4")
![Tile 5](https://raw.githubusercontent.com/adam-szymanski/TilesetSlicer/master/example/tiles/tile_005.png "Tile 5")
![Tile 6](https://raw.githubusercontent.com/adam-szymanski/TilesetSlicer/master/example/tiles/tile_006.png "Tile 6")
![Tile 7](https://raw.githubusercontent.com/adam-szymanski/TilesetSlicer/master/example/tiles/tile_007.png "Tile 7")