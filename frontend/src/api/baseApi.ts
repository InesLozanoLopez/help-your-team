export abstract class BaseApi {
  protected static apiUrl: string = 'http://localhost:8080/api/v1';

  constructor(obj?: BaseApi) {
    obj && Object.assign(this, obj);
  }

  protected static async getUrl(endpoint: string): Promise<string> {
    return `${this.apiUrl}/${endpoint}`;
  }
}
