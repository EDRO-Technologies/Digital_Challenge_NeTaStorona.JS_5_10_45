<script setup lang="ts"></script>

<template>
  <div class="flex gap-6">
    <el-card class="h-[80vh] w-[36vw] flex flex-col gap-4">
      <div class="flex gap-6 mb-4">
        <div class="rounded-[50%] w-[180px] h-[180px] bg-gray-800"></div>
        <div class="flex flex-col pt-12">
          <p class="font-bold text-xl whitespace-break-spaces"
            >Тестов Тест Тестович</p
          >
          <p>Сургут, 38 лет</p>
          <p>front-end разработчик</p>
        </div>
      </div>
      <div class="mb-4">
        <p class="font-bold text-xl">О себе</p>
        <p class="whitespace-break-spaces">
          Lorem ipsum dolor sit amet. Aut molestiae perferendis et pariatur eius
          non rerum dolorum sed alias distinctio? Eos nostrum vero qui
          voluptatem sunt ut aspernatur alias non molestiae doloremque id fugiat
          itaque aut provident aliquam.
        </p>
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
    <div class="w-[60vw]">
      <el-skeleton />
      <el-skeleton />
      <el-skeleton />
      <el-skeleton />
    </div>
  </div>
</template>

<style scoped></style>
