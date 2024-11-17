import axios from "axios";
import { AppEvent } from "@/stores/models/AppEvent";
import { EventGathererServiceAddress } from "@/utils/ServiceAddresses";

export const getEvents = async (): Promise<AppEvent[]> => {
  try {
    const { data } = await axios.get<AppEvent[]>(EventGathererServiceAddress);
    return data.map((event) => new AppEvent(event));
  } catch (e) {
    throw e;
  }
};
