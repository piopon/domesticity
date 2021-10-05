export enum IpcType {
  AddEvent,
  DeleteEvent,
}

export class IpcMessage {
  type: IpcType;
  message: string;

  public constructor(type: IpcType, message: string) {
    this.type = type;
    this.message = message;
  }

  public static newEvent(newEventDate: string): IpcMessage {
    return new IpcMessage(IpcType.AddEvent, newEventDate);
  }

  public static deleteEvent(deletedEventDate: string): IpcMessage {
    return new IpcMessage(IpcType.DeleteEvent, deletedEventDate);
  }
}