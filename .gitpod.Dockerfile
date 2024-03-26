FROM temporalio/gitpod-edu:1.0

USER root

# Add aliases to systemwide profile so they're always available
RUN echo 'alias workspace="cd ${GITPOD_REPO_ROOT}"' >> /etc/profile
RUN echo 'alias webui="gp preview $(gp url 8080)"' >> /etc/profile

# Copy Temporal-specific binaries to a directory we can
# rely on to be in the executable path
RUN cp -L /home/linuxbrew/.linuxbrew/bin/go /usr/local/bin
RUN cp -L /home/linuxbrew/.linuxbrew/bin/gofmt /usr/local/bin
RUN cp -L /home/linuxbrew/.linuxbrew/bin/tctl /usr/local/bin
RUN cp -L /home/linuxbrew/.linuxbrew/bin/tctl-authorization-plugin /usr/local/bin
RUN cp -L /home/linuxbrew/.linuxbrew/bin/temporal /usr/local/bin
