import type { ILegoRequest, ILegoResonse } from "@/stores/models/LegoDto";
import axios from "axios";
import { LlmSendRequestService } from "@/utils/ServiceAddresses";

export const sendLlmRequest = async (
  request: ILegoRequest,
): Promise<ILegoResonse> => {
  try {
    const { data } = await axios.post<ILegoResonse>(
      LlmSendRequestService,
      request,
    );
    return data;
  } catch (e) {
    throw e;
  }
};
