FROM alpine:3.13.6

# Output python straight to terminal
ENV PYTHONUNBUFFERED=1

# Install python3 and make a link to make it the default "python"
RUN apk add --update --no-cache python3 && ln -sf python3 /usr/bin/python

# Ensure pip is installed (Google for more details)
RUN python3 -m ensurepip

# Change to non-root user
RUN adduser -D worker
USER worker
WORKDIR /home/worker

# Add local bin folder to PATH since python stuff is often placed here
RUN mkdir -p /home/worker/.local/bin
ENV PATH="/home/worker/.local/bin:${PATH}"

# Install python setuptools and pip to the latest version
RUN pip3 install --no-cache --upgrade pip setuptools

# Install Robot Framework and some common libs
RUN pip3 install \
		--no-cache-dir \
		requests

COPY sm /home/worker/.local/bin/sm
COPY smlib.py /home/worker/.local/bin/smlib.py

ENTRYPOINT [ "sm" ]