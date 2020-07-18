#!/bin/bash
key_path=$1
source_path=$2
publishto=$3
project=$4
batch_size=$5
products_file=$source_path/products.txt
stores_file=$source_path/stores.txt
echo "$products_file"
echo "$stores_file"
if [ ! -f $products_file ] 
then
   echo "$products_file does not exist." 
   exit 
fi

if [ ! -f $stores_file ] 
then
   echo "$stores_file does not exist." 
   exit 
fi

lines_in_product=$(cat $products_file | wc -l)
noOfbatches=$(( $lines_in_product/$batch_size ))
echo "no of batches $noOfbatches"
for ((i=0;i<$(( $noOfbatches ));i++)); 
do 
name="offersender-$i"
echo "sudo docker run --rm -d -v $key_path:/keys/key.json -v $source_path:/source -e PROJECT=$project -e DESTINATION=$publishto -e BATCHNUMBER=$i -e BATCHSIZE=$batch_size -e CONCURRENCY=$batch_size --n $name lruchandani/offersender"
sudo docker run --rm -d -v $key_path:/keys/key.json -v $source_path:/source -e PROJECT=$project -e DESTINATION=$publishto -e BATCHNUMBER=$i -e BATCHSIZE=$batch_size -e CONCURRENCY=$batch_size --n $name lruchandani/offersender
echo "Started $name"
done