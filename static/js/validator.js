/**
 * Validator
 * 表单验证
 * @type {{isNull: Function, isNotNull: Function, isTelOrPhone: Function, isNotTelAndPhone: Function, isGtLength: Function, isEmail: Function, isNumber: Function}}
 */
var Validator = {

    /**
     *  为NULL
     * @param obj
     * @returns {boolean}
     */
    isNull: function (obj) {
        if (obj == null || obj == undefined) {
            return true;
        }
        if (typeof obj == 'string') {
            return obj.length == 0
        }
        return false;
    },

    /**
     * NOT NULL
     * @param obj
     * @returns {boolean}
     */
    isNotNull: function(obj) {
        return !this.isNull(obj);
    },

    /**
     * 电话or手机号
     * @param obj
     * @returns {boolean}
     */
    isTelOrPhone: function (obj) {
        var regEx = /^1\d{10}$|^(0\d{2,3}-?|\(0\d{2,3}\))?[1-9]\d{4,7}(-\d{1,8})?$/;
        //var regEx = /((\d{11})|^((\d{7,8})|(\d{4}|\d{3})-(\d{7,8})|(\d{4}|\d{3})-(\d{7,8})-(\d{4}|\d{3}|\d{2}|\d{1})|(\d{7,8})-(\d{4}|\d{3}|\d{2}|\d{1}))$)/;
        return regEx.test(obj);
    },

    /**
     * NotTelAndPhone
     * @param obj
     * @returns {boolean}
     */
    isNotTelAndPhone : function(obj) {
        return !this.isTelOrPhone(obj);
    },

    /**
     * 长度 > len
     *
     * @param obj
     * @param len
     * @returns {Boolean}
     */
    isGtLength: function (obj, len) {
        if (this.isNull(obj))
            return false;
        if (obj.length > len) {
            return true;
        }
        return false;
    },

    /**
     * Email验证
     *
     * @param obj
     * @returns
     */
    isEmail: function (str) {
        var regEx = /^[a-z0-9]+([._\\-]*[a-z0-9])*@([a-z0-9]+[-a-z0-9]*[a-z0-9]+.){1,63}[a-z0-9]+$/;
        return regEx.test(str);
    },

    /**
     * 数字验证
     *
     * @param obj
     * @returns
     */
    isNumber: function (obj) {
        var regEx = /^\d*$/;
        return regEx.test(obj);
    }
}









