#!/bin/sh

# membuat base name
mkdir "$1 at $(date)"

cd "$1 at $(date)" 

# membuat folder-folder
mkdir -p "about_me/personal"
mkdir -p "about_me/professional"
mkdir "my_friends"
mkdir "my_system_info"

# create file facebook
file_pos="about_me/personal/facebook.txt"
echo "https://www.facebook.com/$2" > $file_pos

# create file linkedin
file_pos="about_me/professional/linkedin.txt"
echo "https://www.linkedin.com/in/$3" > $file_pos

# create list of friends
file_pos="my_friends/list_of_my_friends.txt"
echo "$(curl -s https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt)" > $file_pos

# create about_this_laptop
file_pos="my_system_info/about_this_laptop.txt"
echo "$1" > $file_pos
echo "$(uname -a)" >> $file_pos

# ping 3 times to google
file_pos="my_system_info/internet_connection.txt"
echo "$(ping -c 3 google.com)" > $file_pos


