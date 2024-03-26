FROM temporalio/gitpod-edu:1.0

# Copy Temporal-specific binaries to a directory we can
# rely on to be in the executable path
COPY /home/linuxbrew/.linuxbrew/Cellar/go/1.22.1/libexec/bin/go /usr/local/bin
COPY /home/linuxbrew/.linuxbrew/Cellar/go/1.22.1/libexec/bin/gofmt /usr/local/bin
COPY /home/linuxbrew/.linuxbrew/Cellar/temporal/0.11.0/bin/temporal /usr/local/bin
COPY /home/linuxbrew/.linuxbrew/Cellar/tctl/1.18.0/bin/tctl /usr/local/bin
#COPY /home/linuxbrew/.linuxbrew/Cellar/tctl/1.18.0/bin/tctl-authorization-plugin /usr/local/bin


# Add bash aliases to systemwide profile so that they'll 
# always be available
COPY .bash_aliases /etc/profile.d

