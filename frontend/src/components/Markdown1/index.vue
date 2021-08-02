<template>
    <div v-html="markdownText" />
</template>

<script>
import axios from 'axios'

export default ({
    data : ()=> {
        return {
            markdownText:"",
            originMarkdownText:"**hello world**"
        }
    },
    created : function() {
        axios.post(
            "https://api.github.com/markdown",
            {
                text: this.originMarkdownText
            },
            {
                headers: {
                    Accept : "application/vnd.github.v3+json"
                }
            }
        ).then(response => {
            this.markdownText = response.data
        }).catch(error => {
            console.error(error)
        })
    }
})
</script>
