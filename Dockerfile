FROM scratch
ADD goapi goapi
ENV PORT 4200
EXPOSE 4200
ENTRYPOINT ["/goapi"]