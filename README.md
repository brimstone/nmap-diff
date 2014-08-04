# nmap-diff
This program augments nmap scans by displaying changes in the network.

# Usage
## Simple Scan Diff
This usage displays the differances between two nmap scans.
`nmap-diff old-scan.xml new-scan.xml`
`nmap-diff -p previous-scan.xml -c current-scan.xml -f html`

## network monitoring
This usage watches network traffic on the specified interface, runs nmap against the host and displays the differances.
`nmap-diff -i wlan0`