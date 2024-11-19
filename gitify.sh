#!/bin/bash
#/////////////

desired_branch="main"
username="SpiXla"
email="aymanaitbihi9@gmail.com"
github_repo="https://github.com/SpiXla/lemin.git"
gitea_repo="https://learn.zone01oujda.ma/git/aaitbihi/lem-in.git"

set -e

# Prints its first argument...
# Adds a 1 second cooldown...
log() {
    echo -e "$1"
    sleep .5
}

# Checks if there is already a username and email in configuration...
# If no username or email is found, it will configure it...
check_credentials() {
    log "\nChecking credentials..."

    if ! git config --get user.name || ! git config --get user.email; then
        log "No Credentials found!!!"

        log "Configuring credentials..."
        git config user.name "$username"
        git config user.email "$email"
    fi

    log "Credentials have been configured successfully!"
}

# Checks if we are in the desired branch before adding files...
# Exits the program if we're in a different branch to prevent confusion...
check_branch() {
    log "\nChecking branch..."
    git branch

    if [ "$(git rev-parse --abbrev-ref HEAD)" != "$desired_branch" ]; then
        log "[WARNING] Not on branch $desired_branch..."
        log "Please restart..."
        exit 1
    fi

    log "You're on branch $desired_branch!"
}

# Adds the remotes for GitHub and Gitea if they are not already set...
configure_remotes() {
    log "\nConfiguring remotes..."

    # Check if 'origin' (GitHub) remote is configured
    if ! git remote get-url origin &>/dev/null; then
        log "GitHub remote not found. Adding GitHub remote..."
        git remote add origin "$github_repo"
    else
        log "GitHub remote already configured."
    fi

    # Check if 'mirror' (Gitea) remote is configured
    if ! git remote get-url mirror &>/dev/null; then
        log "Gitea remote not found. Adding Gitea remote..."
        git remote add mirror "$gitea_repo"
    else
        log "Gitea remote already configured."
    fi

    log "Remotes configured successfully!"
}

# Adds the files given as arguments...
# If there are no arguments, it adds all files...
add_files() {
    log "\nAdding files..."

    if [ $# -eq 0 ]; then
        log "No Files specified..."

        log "Adding all changes..."
        git add .
    else
        log "Adding Specified files..."

        for file in "$@"; do
            git add "$file"
            log "Added: $file"
        done
    fi

    log "Files added successfully!"
}

# Prompts the user for a commit message...
# Then commits the changes...
# Then pushes the changes...
commit_and_push() {
    log "\nTime to push changes..."

    git status
    read -r -p "Enter Commit Message: " commit_message
    git commit -a -m "$commit_message"
    git push origin "$desired_branch"
    git push mirror "$desired_branch"

    log "Well done!"
}

main() {
    check_credentials
    check_branch
    configure_remotes  # Add remotes if not already configured
    add_files "$@"
    commit_and_push
    git remote show origin
}

main "$@"
