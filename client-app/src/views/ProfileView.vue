<template>
  <div class="flex gap-20">
    <el-card class="h-[80vh] w-[36vw] flex flex-col gap-4">
      <div class="flex gap-6 mb-4">
        <div
          class="rounded-[50%] flex justify-center items-center w-[180px] h-[180px] bg-gray-200"
        >
          <el-icon size="150" color="white">
            <UserFilled />
          </el-icon>
        </div>
        <div class="flex flex-col pt-12">
          <p class="font-bold text-xl whitespace-break-spaces"
            >Тестов Тест Тестович</p
          >
          <p>Сургут, 38 лет</p>
          <p>front-end разработчик</p>
        </div>
      </div>

      <div>
        <div class="mb-2">
          <p class="font-bold text-xl">Рейтинг специалиста</p>
          <el-progress
            :text-inside="true"
            :stroke-width="24"
            :percentage="78"
            status="success"
          />
        </div>
        <div>
          <p>JavaScript</p>
          <el-progress
            :text-inside="true"
            :stroke-width="20"
            :percentage="65"
            status="success"
          />
        </div>
        <div>
          <p>C#</p>
          <el-progress
            :text-inside="true"
            :stroke-width="20"
            :percentage="30"
            status="exception"
          />
        </div>
        <div>
          <p>Vue</p>
          <el-progress
            :text-inside="true"
            :stroke-width="20"
            :percentage="45"
            status="warning"
          />
        </div>
      </div>
    </el-card>
    <div class="w-[60vw] flex flex-col gap-4">
      <p class="font-bold text-xl">Опыт</p>
      <div class="flex flex-wrap gap-2">
        <el-tag
          v-for="tag in dynamicTags"
          :key="tag"
          size="large"
          closable
          :disable-transitions="false"
          @close="handleClose(tag)"
        >
          {{ tag }}
        </el-tag>
        <el-input
          v-if="inputVisible"
          ref="InputRef"
          v-model="inputValue"
          class="w-20"
          size="large"
          @keyup.enter="handleInputConfirm"
          @blur="handleInputConfirm"
        />
        <el-button
          v-else
          class="button-new-tag"
          size="default"
          @click="showInput"
        >
          + New Tag
        </el-button>
      </div>
      <el-upload
        class="upload-demo"
        drag
        action="https://run.mocky.io/v3/9d059bf9-4660-45f2-925d-ce80ad6c4d15"
        multiple
      >
        <el-icon class="el-icon--upload">
          <upload-filled />
        </el-icon>
        <div class="el-upload__text h-[12vh]">
          Перетащите сюда или <em>Нажмите чтобы загрузить сертификат</em>
        </div>
        <template #tip>
          <div class="el-upload__tip">
            размер файлов должен быть меньше 5мб
          </div>
        </template>
      </el-upload>
      <el-skeleton />
    </div>
  </div>
</template>

<script setup lang="ts">
import { nextTick, ref } from "vue";
import type { InputInstance } from "element-plus";
import { UploadFilled, UserFilled } from "@element-plus/icons-vue";

const dynamicTags = ref(["бэкенд", "девопс", "машин лернинг", "фронтенд"]);
const inputValue = ref("");
const inputVisible = ref(false);
const InputRef = ref<InputInstance>();

const handleClose = (tag: string) => {
  dynamicTags.value.splice(dynamicTags.value.indexOf(tag), 1);
};

const showInput = () => {
  inputVisible.value = true;
  nextTick(() => {
    InputRef.value!.input!.focus();
  });
};

const handleInputConfirm = () => {
  if (inputValue.value) {
    dynamicTags.value.push(inputValue.value);
  }
  inputVisible.value = false;
  inputValue.value = "";
};
</script>

<style scoped></style>
