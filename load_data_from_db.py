import os
import sqlite3

from PIL import Image, ImageDraw


def midPoint(xmin, xmax, ymin, ymax):
    return float(xmin + xmax) * 0.5, float(ymin + ymax) * 0.5


db_path = './image_processing.db'
save_img_directory = './images/'

connection = sqlite3.connect(db_path)
cursor = connection.cursor()

img_path = "shop/youtube_shop0"
img_number = 86
max_img_number = 376
img_step = 10

colors = ['white', 'pink', 'yellow', 'red', 'blue', 'green']
map_id_object_color = dict()
map_id_object_mid_points = dict()
outline_width = 5
color_counter = 0

if not os.path.exists(save_img_directory):
    os.makedirs(save_img_directory)

while img_number <= max_img_number:
    if img_number < 100:
        str_img_number = '0' + str(img_number)
        img_file = img_path + str_img_number + '.jpg'
    else:
        str_img_number = str(img_number)
        img_file = img_path + str_img_number + '.jpg'

    cursor.execute("SELECT * FROM objects WHERE img_id = " + str(img_number))
    result = cursor.fetchall()
    im = Image.open(img_file)
    # create rectangle image
    img1 = ImageDraw.Draw(im)

    for r in result:
        x_min, x_max = r[3], r[4]
        y_min, y_max = r[5], r[6]
        object_id = r[2]
        midX, midY = midPoint(x_min, x_max, y_min, y_max)
        if object_id not in map_id_object_mid_points:
            map_id_object_mid_points[object_id] = [(midX, midY, midX + 4, midY + 4)]
        else:
            map_id_object_mid_points[object_id].append((midX, midY, midX + 4, midY + 4))

        if object_id not in map_id_object_color:
            map_id_object_color[object_id] = colors[color_counter % len(colors)]
            color_counter += 1
        color = map_id_object_color[object_id]

        shape = [(x_min, y_min), (x_max, y_max)]
        img1.rectangle(shape, outline=color, width=outline_width)
        points = map_id_object_mid_points[object_id]
        if len(points) > 10:
            points = points[len(points) - 10:]
        for point in points:
            img1.ellipse(point, fill=color)
        text_shape = (x_min, y_min)
        fill = (x_min + 120, y_min + 10)
        img1.rectangle([text_shape, fill], fill=color)
        img1.text(text_shape, object_id, fill='black')

    # write to stdout
    im.save(save_img_directory + "img" + str_img_number + '.jpg', "JPEG")
    img_number += img_step
