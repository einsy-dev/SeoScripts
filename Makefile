#!make
include .env

# The 'all' target is the default. It depends on 'deploy' in this example.
all: 
	gow

# The 'sync' target copies local files to the remote server
sync:
	wsl rsync $(RSYNC_OPTS) $(LOCAL_DIR) $(REMOTE_USER)@$(REMOTE_HOST):$(REMOTE_DIR)

# The 'deploy' target first syncs files, then executes a command on the remote server via SSH
deploy: sync
	wsl ssh -t ${REMOTE_USER}@$(REMOTE_HOST) "cd ~/app && sudo docker compose up -d --build --force-recreate"

backup:
	wsl rsync -avz --rsync-path="sudo rsync" $(REMOTE_USER)@$(REMOTE_HOST):~/app/volumes/ /mnt/e/Backups/SeoDatabase/$$(date +%Y.%m.%d.%H.%M.%S)
# A 'dry-run' target to test the rsync command without making actual changes
dry-run:
	rsync $(RSYNC_OPTS) -n $(LOCAL_DIR) $(REMOTE_USER)@$(REMOTE_HOST):$(REMOTE_DIR)

.PHONY: all deploy sync dry-run