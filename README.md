# go-sync-deleted-raw-jpg

This project deletes RAW images after JPGs have been sorted out, as sorting JPGs is much faster then with RAW files, at least in MAC OSX Finder.
So my image editing process is as follows:

* Shoot both RAW and JPG images
* Split RAW and JPG images into a `JPG` and a `RAW` directory (even though I'm a sony user, so the raw images do have an `ARW` file extension)
* Sort JPGs, delete crappy images
* run this tool to delete the RAW files from which I deleted the JPGs

# License

```
MIT License

Copyright (c) 2020 Alex Klinkert
```
