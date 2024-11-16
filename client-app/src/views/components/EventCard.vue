<template>
  <div class="flex h-[26vh]">
    <div
      class="font-bold text-xl flex flex-col items-center justify-center bg-cyan-800 text-white px-1 -mr-0.5"
    >
      <p>17</p>
      <p>Ноября</p>
    </div>
    <img
      :src="props.source"
      alt="Изображение мероприятия"
      width="200"
      height="200"
    />
    <div class="flex flex-col gap-1 ml-2">
      <p class="font-bold text-xl">{{ props.title }}</p>
      <p>
        Lorem ipsum dolor sit amet. Est blanditiis excepturi sit velit
        praesentium ut voluptate quis ut obcaecati mollitia sit nobis sunt nam
        explicabo voluptatem. Qui labore voluptas non tempore quod est dolores
        repellendus? Est consequatur voluptatum ut dolores dolores et beatae
        quasi id maiores dolor aut perferendis dolor. Quo inventore facilis eos
        amet magni ut delectus libero. Ut voluptatem aliquid qui repellendus
        reiciendis et nesciunt recusandae ex esse delectus ut quidem voluptate.
      </p>
      <el-button type="primary" class="mt-8 self-end max-w-40"
        >Подробнее
      </el-button>
    </div>
  </div>
</template>

<script setup lang="ts">
const props = defineProps<{
  source: string;
  title: string;
  description: string;
}>();
</script>

<style scoped></style>
