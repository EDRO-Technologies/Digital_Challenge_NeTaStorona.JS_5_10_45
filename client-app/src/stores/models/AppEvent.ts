export interface IAppEvent {
  id: number;
  full_name: string;
  full_info: string;
  date_start: Date;
  date_end: Date;
  themes: {
    name: string;
  };
  type: {
    name: string;
  };
  photo: string;
  format: string;
  space: {
    name: string;
    address: {
      city: string;
      title: string;
    };
  } | null;
}

export class AppEvent implements IAppEvent {
  constructor(
    options: IAppEvent = {
      id: 0,
      full_name: "",
      full_info: "",
      date_start: new Date(),
      date_end: new Date(),
      themes: { name: "" },
      type: { name: "" },
      photo: "/Pattern.png",
      format: "",
      space: null,
    },
  ) {
    this.id = options.id;
    this.full_name = options.full_name;
    this.full_info = this.parseFullInfo(options.full_info);
    this.date_start = new Date(options.date_start);
    this.date_end = new Date(options.date_end);
    this.themes = options.themes;
    this.type = options.type;
    this.photo = options.photo;
    this.format = options.format;
    this.space = options.space;
  }

  id: number;
  full_name: string;
  full_info: string;
  date_start: Date;
  date_end: Date;
  themes: {
    name: string;
  };
  type: {
    name: string;
  };
  photo: string;
  format: string;
  space: {
    name: string;
    address: {
      city: string;
      title: string;
    };
  } | null;

  parseFullInfo(data: string): string {
    try {
      const parsed = JSON.parse(data);
      return parsed.blocks.map((block: any) => block.data.text).join(" ");
    } catch (e) {
      return data;
    }
  }
}
