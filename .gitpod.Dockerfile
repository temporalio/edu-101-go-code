FROM temporalio/gitpod-edu:1.0

USER root

# Customize the shell for the gitpod user, which may not yet exist
# when this runs, so we achieve that by modifying the bash RC file
# that's copied to a new user's home directory. 
RUN echo 'alias workspace="cd ${GITPOD_REPO_ROOT}"' >> /etc/skel/.bashrc
RUN echo 'alias webui="gp preview $(gp url 8080)"' >> /etc/skel/.bashrc

# These will provide a shorter, but still useful, shell prompt
RUN echo 'export PROMPT_DIRTRIM=2' >> /etc/skel/.bashrc
RUN echo 'export PS1="\[\033[01;34m\]\w\[\033[00m\]\$ "' >> /etc/skel/.bashrc


# Copy Temporal-specific binaries to a directory we can
# rely on to be in the executable path even after a timeout
RUN cp -L /home/linuxbrew/.linuxbrew/bin/go /usr/local/bin
RUN cp -L /home/linuxbrew/.linuxbrew/bin/gofmt /usr/local/bin
RUN cp -L /home/linuxbrew/.linuxbrew/bin/tctl /usr/local/bin
RUN cp -L /home/linuxbrew/.linuxbrew/bin/tctl-authorization-plugin /usr/local/bin
RUN cp -L /home/linuxbrew/.linuxbrew/bin/temporal /usr/local/bin
