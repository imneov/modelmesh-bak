file_list=( $(find {api,mindspore_serving} -name '*.proto') )

for file in "${file_list[@]}"; do
    echo "Processing ${file}"
done
