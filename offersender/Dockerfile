FROM  alpine
COPY ./bin ./product.template ./offer.template  /
VOLUME [ "/keys","/source" ]
ENV GOOGLE_APPLICATION_CREDENTIALS=/keys/key.json
CMD /offersender -project-id=$PROJECT -publish-to=$DESTINATION -batch-number=$BATCHNUMBER -batch-size=$BATCHSIZE  -concurrent-requests=$CONCURRENCY -source-file-path=/source