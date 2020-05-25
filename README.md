# Scanning-Live-Subdomain ( Using Dig Command )
Take a list of domains and filter out outdated subdomain

# Install
git clone https://github.com/ManShum815/Automating-Scan-Live-Subdomain or just copy my script from live folder

# Setup
cd Scanning-Live-Subdomain 

cd live

chmod +x scan.sh

./scan.sh

# Script 
while read target; do
if dig "$target" > /dev/null; then
# The subdomain is outdated, then do not
# put into a file called "target.txt".
echo "$target" >> target.txt
fi
done < /path/of/(your-subdomain-list)
