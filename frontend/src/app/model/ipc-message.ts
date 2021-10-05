export enum IpcType {
  AddEvent,
  DelEvent,
}

export class IpcMessage {
  type: IpcType;
  message: string;

  public constructor(type: IpcType, message: string) {
    this.type = type;
    this.message = message;
  }

  public static newEvent(message: string): IpcMessage {
    return new IpcMessage(IpcType.AddEvent, message);
  }
}
