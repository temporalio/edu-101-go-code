FROM temporalio/gitpod-edu:1.0

# Copy Temporal-specific binaries to a directory we can
# rely on to be in the executable path
RUN /usr/bin/cp -L /home/linuxbrew/.linuxbrew/bin/* /usr/local/bin

# Append bash aliases to systemwide profile so that 
# they'll always be available
RUN cat .bash_aliases >> /etc/profile

