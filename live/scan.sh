while read target; do
if host "$target" > /dev/null; then
# The subdomain is outdated, then do not
# put into a file called "target.txt".
echo "$target" >> target.txt
fi
done < /path/of/(your-subdomain-list)
