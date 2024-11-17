<template>
  <div
    v-loading.fullscreen.lock="isLoading"
    class="quiz flex flex-col h-[80vh] mx-auto"
  >
    <el-input
      v-model="requestMessage"
      placeholder="Введите ваш запрос"
      class="mb-4"
    />
    <div
      class="pagination flex justify-between items-center p-4 border-b border-gray-300 bg-gray-50"
    >
      <el-button @click="prevPage" :disabled="currentPage === 0">
        Назад
      </el-button>
      <span>Вопрос {{ currentPage + 1 }} из {{ questions.length }}</span>
      <el-button @click="nextPage">Вперёд</el-button>
    </div>
    <div class="content flex-grow p-4 overflow-y-auto">
      <h2 class="text-xl font-semibold mb-4">{{ currentQuestion.question }}</h2>
      <div
        v-for="(answer, index) in currentQuestion.answers"
        :key="index"
        class="answer-option mb-4"
      >
        <label class="flex items-center space-x-2">
          <input
            type="checkbox"
            :value="answer"
            :checked="isAnswerSelected(index)"
            @change="toggleAnswer(index)"
            class="form-checkbox bg-cyan-800"
          />
          <span>{{ answer }}</span>
        </label>
      </div>
    </div>
    <div
      v-if="currentPage === questions.length - 1"
      class="submit-container flex justify-center mt-8"
    >
      <el-button
        type="primary"
        size="large"
        @click="submitAnswers"
        class="w-full"
        :disabled="requestMessage.length === 0"
        >Отправить
      </el-button>
    </div>
  </div>
  <ModelAnswer :answer="marked(answer)" />
</template>

<script lang="ts">
import { surveyQuestions } from "@/utils/SurveyQuestions.js";
import { sendLlmRequest } from "@/services/LegoService.js";
import type { ILegoRequest } from "@/stores/models/LegoDto";
import { marked } from "marked";
import ModelAnswer from "@/views/components/ModelAnswer.vue";

export default {
  components: { ModelAnswer },
  data() {
    return {
      questions: surveyQuestions,
      currentPage: 0,
      selectedAnswers: {} as Record<number, string[]>,
      requestMessage: "",
      answer: "",
      isLoading: false,
    };
  },
  computed: {
    currentQuestion() {
      return this.questions[this.currentPage];
    },
  },
  methods: {
    marked,
    isAnswerSelected(answerIndex: number) {
      return (
        this.selectedAnswers[this.currentPage]?.includes(
          `${this.currentQuestion.question}: ${this.currentQuestion.answers[answerIndex]}`,
        ) || false
      );
    },
    toggleAnswer(answerIndex: number) {
      const selected =
        this.selectedAnswers[this.currentPage] || ([] as string[]);
      const data = `${this.currentQuestion.question}: ${this.currentQuestion.answers[answerIndex]}`;
      if (selected.includes(data)) {
        this.selectedAnswers[this.currentPage] = selected.filter(
          (item: any) => item !== data,
        );
      } else {
        this.selectedAnswers[this.currentPage] = [...selected, data];
      }
    },
    prevPage() {
      if (this.currentPage > 0) {
        this.currentPage--;
      }
    },
    nextPage() {
      if (this.currentPage < this.questions.length - 1) {
        this.currentPage++;
      }
    },
    prepareRequest() {
      const context = Object.values(this.selectedAnswers).flat().join("; ");
      return {
        context: context,
        message: `${this.requestMessage}. Исходя из результатов опросника, составь роадмап к достижению моей цели`,
      } as ILegoRequest;
    },
    async submitAnswers() {
      const data = this.prepareRequest();
      try {
        this.isLoading = true;
        this.answer = await sendLlmRequest(data);
        this.isLoading = false;
      } catch (error) {
        this.isLoading = false;
      }
    },
  },
};
</script>

<style scoped></style>
