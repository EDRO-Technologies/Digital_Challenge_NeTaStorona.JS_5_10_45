import { defineStore } from "pinia";
import { AppEvent } from "@/stores/models/AppEvent";
import { computed, ref } from "vue";

export const useEventStore = defineStore("EventStore", () => {
  const appEventsState = ref<AppEvent[]>([]);

  function $reset() {
    appEventsState.value = [];
  }

  const appEvents = computed(() =>
    appEventsState.value.map((ev) => new AppEvent(ev)),
  );

  const storeEvents = (events: AppEvent[]) => {
    appEventsState.value = events;
  };

  return {
    $reset,
    appEventsState,
    appEvents,
    storeEvents,
  };
});
