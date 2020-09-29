<template>
    <div class="content">
        <div class="message-content">
            <div 
            v-for="(item, key) in data" 
            class="message" :class="{'me': item.source == 'me', 'system': item.source != 'me'}"
            :key="key"
            >
                <span>{{item.content}}</span>
            </div>
        </div>
        <div class="form">
            <form @submit.prevent.stop="submit">
                <input type="text" v-model="message"/>
                <button type="submit">提交</button>
            </form>
        </div>
    </div>
</template>
<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'

interface IMesage {
    source: string
    content: string
}

@Component
export default class Home extends Vue {

    data: Array<IMesage> = []

    message: string = ''

    submit() {
        const data: IMesage = {
            source: "me",
            content: this.message,
        } 

        this.message = ''
        this.data.push(data)
    }
}
</script>

<style>
    .content  {
        background: red;
    }
    .me {
        color: green;
        text-align: right;
    }

    .system {
        color: orange;
        text-align: left;
    }

    .message {
        padding: 5px;
    }


    .form {
        flex: 1 1 auto;
        height: 35px;
    }
</style>