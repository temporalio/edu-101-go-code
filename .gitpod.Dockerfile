FROM temporalio/gitpod-edu:1.0

USER root

# Add aliases to systemwide profile so they're always available
RUN echo 'alias workspace="cd ${GITPOD_REPO_ROOT}"' >> /etc/profile
RUN echo 'alias webui="gp preview $(gp url 8080)"' >> /etc/profile

# Customize the shell prompt (specifically for the gitpod user, which
# may not exist when this runs, so we customize the bash RC file that
# gets copied to a new user's home directory). The goal of these 
# customizations is a shorter, but still useful, shell prompt.
RUN echo 'export PROMPT_DIRTRIM=2' >> /etc/skel/.bashrc
RUN echo 'export PS1="\[\033[01;34m\]\w\[\033[00m\]\$ "' >> /etc/skel/.bashrc


# Copy Temporal-specific binaries to a directory we can
# rely on to be in the executable path even after a timeout
RUN cp -L /home/linuxbrew/.linuxbrew/bin/go /usr/local/bin
RUN cp -L /home/linuxbrew/.linuxbrew/bin/gofmt /usr/local/bin
RUN cp -L /home/linuxbrew/.linuxbrew/bin/tctl /usr/local/bin
RUN cp -L /home/linuxbrew/.linuxbrew/bin/tctl-authorization-plugin /usr/local/bin
RUN cp -L /home/linuxbrew/.linuxbrew/bin/temporal /usr/local/bin
