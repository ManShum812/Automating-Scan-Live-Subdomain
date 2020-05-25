// Copyright © 2020 Man Shum

package main

while read target; do
if dig "$target" > /dev/null; then
# The subdomain is outdated, then do not
# put into a file called "target.txt".
echo "$target" >> target.txt
fi
done < /path/of/(your-subdomain-list)
