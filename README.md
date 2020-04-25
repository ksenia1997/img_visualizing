
##Storing data in Golang
- Sequentially load xml files with labels – they include bounding boxes of detected objects
- Store the labels into a database 
    - They have to include coordinates and ID of a bounding box (and maybe something more if necessary)
        -  In the first picture, you assign the IDs randomly
        -  In the following ones, you assign the ID of the closest bounding box from the previous picture (it is kind of a simple tracking)

```cassandraql
go build
./img_visualizing
```

xml files are loaded from directory *./shop* (I suppose that imgs numbers are from 86 to 376 with a difference of 10). 
Parsed objects are save in the local database *image_processing.db* in the format: 

- **id** as a Primary Key
- **img_id** is a number of image
- **object_id** is a unique object id with help of https://github.com/rs/xid
- **x_min** bounding box coordinates
- **x_max** bounding box coordinates
- **y_min** bounding box coordinates
- **y_max** bounding box coordinates

## Visualizing in Python

- Sequentially load the images, find their corresponding labels in the database and draw them into the picture
- Use different colours for different IDs of a bounding box
- Show the ID number on top of the bounding box
- With the same colour, draw also the last 10 locations of the central point of the bounding box (it will show how the person moves)
- Save the whole visualization as a video with a framerate 5fps

```cassandraql
python3 load_data_from_db.py
```

Data loaded from *image_processing.db* database. Images with painted borders are saved in *./images*  


To save images as video in the directory ./images:
```cassandraql
ffmpeg -framerate 5 -pattern_type glob -i '*.jpg' video.avi
```
