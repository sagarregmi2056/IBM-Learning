 #!/bin/sh

# Secret message
secret="Lxwpajcdujcrxwb! Hxd dwblajvkunm cqn vnbbjpn rw cqn dbb lqjuunwpn!!"

# Function to rotate lowercase letters
rotate() {
  echo "$1" | tr 'a-z' "$(echo {a..z} | tr -d ' ' | cut -c$2-26)$(echo {a..z} | tr -d ' ' | cut -c1-$2)"
}

# Loop through all possible shifts (1-26)
i=1
while [ $i -le 26 ]; do
  echo "Shift $i:"
  echo "$secret" | tr 'a-z' $(echo {a..z} | tr -d ' ' | cut -c$((i+1))-26)$(echo {a..z} | tr -d ' ' | cut -c1-$i)
  echo "-----------------------"
  i=$((i + 1))
done




The given bash script will find the steps(round) to get the actual answers 