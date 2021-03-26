rootfs:
	docker build -t rootfs . && \
	id=$(docker create rootfs true) && \
	mkdir -p rootfs && \
	docker export "${id}" | tar -x -C rootfs && \
	docker rm -vf "${id}" && \
	docker rmi rootfs

create:
	docker plugin create eslizn/cosfs ./


