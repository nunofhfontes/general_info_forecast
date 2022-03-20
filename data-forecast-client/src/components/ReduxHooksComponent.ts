import React, {ChangeEvent, FormEvent, useState, useEffect} from "react";
import {Form, Input, Button} from "antd";
 
interface Props {
}

const ReduxHooksComponent: React.FC<Props> = () => {
    return (
        <Form layout="inline">
            <Form.Item>
                <Input type="text" placeholder="name"/>
                <Input type="text" placeholder="address" />
                <Button htmlType="submit" type="primary"> Submit </Button>
            </Form.Item>
        </Form>
    )
};

export default ReduxHooksComponent;