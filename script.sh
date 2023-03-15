#!/bin/bash

function replace_with_lando {
  # Replace "idk" with "lando" and run the command
  echo "Running as lando."
  eval "lando ${@/'idk'/''}"
}

function replace_with_ddev {
  # Replace "idk" with "ddev" and run the command
  echo "Running as ddev."
  eval "ddev ${@/'idk'/''}"
}

# Check which development tool config file are present.
if [ -f "./.lando.yml" ] && [ -d "./.ddev" ]
then
  echo -e "Both .lando.yml and .ddev directories exist. \nWhich one would you like to use? Type 'lando' or 'ddev':"
  read option
  case "$option" in
    "lando") 
      replace_with_lando "$@"
      ;;
    "ddev") 
      replace_with_ddev "$@"
      ;;
    *) 
      echo "Invalid option. Exiting."
      exit 1
      ;;
  esac
elif [ -d "./.ddev" ]
then
  replace_with_ddev "$@"
elif [ -f "./.lando.yml" ]
then
  replace_with_lando "$@"
else
  echo "Woopie!! no ddev no lando."
fi
