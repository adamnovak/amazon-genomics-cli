FROM public.ecr.aws/amazonlinux/amazonlinux:2 AS final

# COPY THIRD-PARTY /opt/
COPY LICENSE /opt/

RUN yum update -y \
    && yum install -y \
    curl \
    hostname \
    "java-11-amazon-corretto-headless(x86-64)" \
    unzip \
    jq \
    && yum clean -y all \
    && rm -rf /var/cache/yum

# install awscli v2
RUN curl -s "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "/tmp/awscliv2.zip" \
    && unzip -q /tmp/awscliv2.zip -d /tmp \
    && /tmp/aws/install -b /usr/bin \
    && rm -rf /tmp/aws*

##### MODIFY #######
## In this area install your new engine into the container as well as any requirements for that engine.
## Dockerfile documentation is found here: https://docs.docker.com/engine/reference/builder/

RUN yum update -y \
    && yum install -y \
    python3 \
    && yum clean -y all \
    && rm -rf /var/cache/yum

COPY THIRD-PARTY /opt/
ARG TOIL_VERSION="issues/3942-aws-batch-batch-system"
ARG TOIL_DOWNLOAD_NONCE="076d99707d3cf9c8fcd763e2a7fa36bbc873b8d9"
RUN python3 -m pip install https://github.com/DataBiosphere/toil/archive/${TOIL_VERSION}.zip?${TOIL_DOWNLOAD_NONCE}#egg=toil[aws,cwl,server]

# copy the entrypoint script to the image
COPY toil.aws.sh /opt/bin/toil.aws.sh
RUN chmod +x /opt/bin/toil.aws.sh

#### END MODIFY ######

WORKDIR /opt/work
ENTRYPOINT ["/opt/bin/toil.aws.sh"]

