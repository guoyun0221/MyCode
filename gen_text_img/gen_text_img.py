from PIL import Image, ImageDraw, ImageFont

# create img
width, height = 1920, 1080
image = Image.new("RGB", (width, height), (255, 255, 255)) 
# create a draw object
draw = ImageDraw.Draw(image)

# gradual background color
for y in range(height):
    r = int(224 + (255 - 224) * (y / height))  
    g = int(255 - (255 - 240) * (y / height))  
    b = int(255 - (255 - 225) * (y / height))  
    for x in range(width):
        draw.point((x, y), fill=(r, g, b))

# add text in center of image
font = ImageFont.truetype("simsun.ttc", 140, encoding="unic")
text = "This is text"
text_width, text_height = draw.textsize(text, font)
x = (width - text_width) // 2
y = (height - text_height) // 2
draw.text((x, y), text, fill=(0, 0, 0), font=font)  

# save as png file
image.save("generated_img.png")

# show it
image.show()

