import type { ILegoRequest, ILegoResonse } from "@/stores/models/LegoDto";
import axios from "axios";

const url = "http://172.29.144.204:70/api/prompt";

export const sendLlmRequest = async (
  request: ILegoRequest,
): Promise<ILegoResonse> => {
  try {
    const response = axios.post<ILegoResonse>(url, request);
    console.log(response);
    return (await response).data;
  } catch (e) {
    throw e;
  }
};
