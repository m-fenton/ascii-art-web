ascii-art-web

This project uses Golang's HTML/Templates & net/HTTP packages to create an API that creates a static server & listens and responds to an accompanying HTML file.
The primary aim is take a string from the user and return it to the client in Ascii form.
The user also has the added option of downloading the resulting artwork by exporting a txt file.


Usage: How to run

    
1.  Within the project path: "go run .", to run the server. 
2.  In the web browser, go to localhost:8080
3.  Here the user can chose between 3 banners, shadow, standard and thinkertoy.
4.  In the text area below, the user can input up to 2 lines.
5.  After pressing submit, the ascii art version is returned with the banner chosen.
6.  To download the result press download.

Created by

Martin Fenton
Rupert Cheetham
Nikoi Kwashie