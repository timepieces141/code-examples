import re
from os import listdir
from canistreamit import search, streaming, xfinity

regex = re.compile("^(.*)\(.*$")
for f in listdir("/media/epeters/Seagate Backup Plus Drive/MacGyver_1/TiVo Recordings"):
    match = regex.match(f)
    if match:
        search_list = search(match.group(1))
        for movie in search_list:
            print movie['title']
            stream_result = streaming(movie['_id'])
            if "amazon_prime_instant_video" in stream_result:
                print "\tAmazon Instant"
            if "netflix_instant" in stream_result:
                print "\tNetflix Instant"
            xfinity_result = xfinity(movie['_id'])
            if "streampix" in xfinity_result:
                print "\tXfinity Streampix"
